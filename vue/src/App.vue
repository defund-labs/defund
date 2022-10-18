<template>
  <div>
    <SpTheme>
      <Navbar
        :links="navbarLinks"
        :active-route="router.currentRoute.value.path"
      />
      <router-view />
    </SpTheme>
    <div v-if="store.showTxStatus" class="tx-status-div">
      <Sending v-if="store.sendingTx"></Sending>
      <Success v-if="store.showTxSuccess">
        <a :href="'https://defund.explorers.guru/transaction/' + store.lastTxHash" target="_blank" style="text-decoration: none;">View Transaction</a>
      </Success>
      <Warning v-if="store.showTxFail">
        <a :href="'https://defund.explorers.guru/transaction/' + store.lastTxHash" target="_blank" style="text-decoration: none;">View Transaction</a>
        <p>{{store.lastTxLog}}</p>
      </Warning>
    </div>
  </div>
</template>

<script lang="ts">
import { SpTheme } from '@starport/vue'
import Navbar from './components/Navbar.vue'
import { computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import Success from './components/Success.vue'
import Warning from './components/Warning.vue'
import Sending from './components/Sending.vue'
import { store } from './store/local/store.js'
import { initializeApp } from "firebase/app";
import flatten from 'flat';
import _ from 'lodash';
export default {
  components: { SpTheme, Navbar, Success, Warning, Sending },

  beforeCreate() {
    // Setup Firebase
    const firebaseConfig = {
      apiKey: "AIzaSyCAEhEsltgtOuYI7cEzCEaRbk2ivjb3ucQ",
      authDomain: "defund-testnet.firebaseapp.com",
      projectId: "defund-testnet",
      storageBucket: "defund-testnet.appspot.com",
      messagingSenderId: "504104480253",
      appId: "1:504104480253:web:d6d23ef58070bdb9757d94",
      measurementId: "G-P5VM8QDYBY"
    };

    // Initialize Firebase
    initializeApp(firebaseConfig);
  },

  setup() {

    // store
    let $s = useStore()

    // router
    let router = useRouter()

    // state
    let navbarLinks = [
      { name: 'Portfolio', url: '/portfolio' },
      { name: 'Funds', url: '/funds' },
      { name: 'Stake', url: '/stake' },
      { name: 'Gov', url: '/gov' },
      { name: 'Faucet', url: '/faucet' }
    ]

    // computed
    let address = computed(() => $s.getters['common/wallet/address'])

    // lh
    onBeforeMount(async () => {
      await $s.dispatch('common/env/init')

      // create list of where we will store validators
      var validators = []
      // init the stores we need
      var validators_init = await $s.dispatch("cosmos.staking.v1beta1/QueryValidators", {subscribe: true, all: false})
      // calculate the number of pages of validators
      var numberValPages = Math.ceil(Number(validators_init.pagination.total)/100)
      // dispatch calls for all validators
      var validators_raw = [...Array(numberValPages).keys()].map(async (page) => {
        var new_validators = await $s.dispatch("cosmos.staking.v1beta1/QueryValidators", {query: { "pagination.offset": page * 100 }, subscribe: false, all: false})
        return new_validators
      })
      // resolve all pages of validators
      var valsPromises = await Promise.all(validators_raw)
      // map the validators to a full list of all vals from a list of pages
      var vals = valsPromises.map(async (val) => {
        var new_val = JSON.parse(JSON.stringify(val["validators"]))
        validators = [...validators, ...new_val]
        return JSON.parse(JSON.stringify(val["validators"]))
      })
      // map each validator to flatten object
      var flatVals = validators.map(async (val) => {
          // convert each percentage into a number % (* 100)
          val["commission"]["commission_rates"]["max_change_rate"] = _.round(Number(val["commission"]["commission_rates"]["max_change_rate"]) * 100, 2)
          val["commission"]["commission_rates"]["max_rate"] = _.round(Number(val["commission"]["commission_rates"]["max_rate"]) * 100, 2)
          val["commission"]["commission_rates"]["rate"] = _.round(Number(val["commission"]["commission_rates"]["rate"]) * 100, 2)
          val["commission"]["commission_rates"]["rate_string"] = String(val["commission"]["commission_rates"]["rate"]) + "%"
          // convert all tokens to FETF from ufetf
          val["min_self_delegation"] = _.round(Number(val["min_self_delegation"])/1000000, 2)
          val["tokens"] = _.round(Number(val["tokens"])/1000000, 2)
          val["tokens_string"] = String(val["tokens"]) + " FETF"
          return flatten(val)
      })

      var finalValsResolved = await Promise.all(flatVals)

      var finalVals= _.orderBy(finalValsResolved, (o) => {
        return +o.tokens
      }, ["desc"])

      // resolve all promises for all the flattened validators
      store.validators = finalVals
    })

    return {
      navbarLinks,
      // router
      router,
      // computed
      address,
      store,
      validators: {}
    }
  }
}
</script>

<style scoped lang="scss">
body {
  margin: 0;
  overflow: hidden;
}
.tx-status-div {
  z-index: 500;
  position: fixed;
  bottom: 10px;
  width: 40%;
  border-radius: 10px;
  right: 10px;
  background-color: white;
  padding: 15px;
  border-style: outset;
  border-color: rgba(128, 128, 128, 0.055);
  border-width: 2px;
}
</style>
