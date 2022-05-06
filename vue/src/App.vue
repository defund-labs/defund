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
import { store } from './store/local/popup'
export default {
  components: { SpTheme, Navbar, Success, Warning },

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
    })

    return {
      navbarLinks,
      // router
      router,
      // computed
      address,
      store
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
