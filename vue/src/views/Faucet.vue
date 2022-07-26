<template>
  <div>
    <div v-if="store.twitter == null" class="faucet-div" style="height:70vh;" align="center">
      <button style="background-color:#067acc; border-color:#067acc" @click="login" class="sp-button">Verify Twitter</button>
    </div>
    <div v-else class="faucet-div" style="height:60vh;" align="center">
      <h4 v-if="eligible" style="font-size:x-large; font-weight: bold; margin-bottom: 25px;">You are eligible for the private testnet!</h4>
      <h4 v-if="!eligible" style="font-size:x-large; font-weight: bold; margin-bottom: 25px;">You are not eligible for the private testnet.</h4>
      <button v-if="eligible" :disabled="!eligible ||	airdropped || clicked" @click="requestTokens" style="background-color:#6CE5E8; border-color:#6CE5E8" class="sp-button">{{airdropped ? "Already Airdropped" : "Request Tokens"}}</button>
    </div>
  </div>
</template>

<script>
import { getAuth, signInWithPopup, TwitterAuthProvider } from "firebase/auth";
import { store } from '../store/local/store.js';
import { computed } from 'vue';
import { useStore } from 'vuex';
export default {
  name: 'Faucet',
  async mounted() {
    await this.checkEligibility()
  },
  async onRenderTriggered() {
    await this.checkEligibility()
  },
  data() {
    const $s = useStore()
    var eligible
    var airdropped

    const provider = new TwitterAuthProvider();
    const auth = getAuth();
    let address = computed(() => {
        return $s.getters['common/wallet/address']
    })

    if (store.twitter) {
      const response = this.axios.get(process.env.VUE_APP_DEFUND_USERS_API + store.twitter.providerData[0].uid)
      if(response.data) {
        eligible = response.data.eligible
        airdropped = response.data.data.airdropped
      }
    }

    return {
      auth,
      provider,
      store,
      eligible: eligible ? eligible : false,
      airdropped: airdropped ? airdropped : false,
      address,
      clicked: false
    }
  },
  methods:{
    async login() {
      const result = await signInWithPopup(this.auth, this.provider)
      const user = await result.user;
      store.twitter = user
      this.checkEligibility()
    },
    async checkEligibility() {
      if (store.twitter) {
        const response = await this.axios.get(process.env.VUE_APP_DEFUND_USERS_API + store.twitter.providerData[0].uid)
        this.eligible = await response.data.eligible
        this.airdropped = await response.data.data.airdropped
        return response
      }
      return null
    },
    async markUserAsAirdropped() {
      const response = await this.axios.put(process.env.VUE_APP_DEFUND_USERS_API, { "id": store.twitter.providerData[0].uid })
      return response
    },
    async requestTokens() {
      this.clicked = true
      const res = await this.checkEligibility()
      const eligible = await res.data.eligible
      if(eligible) {
        const ret = await this.axios.post('https://cors.defund.app/' + process.env.VUE_APP_DEFUND_FAUCET_API, {
          address: this.address,
          coins: [
            "100000000ufetf"
          ]
        })
        if (ret.statusCode == 200) {
          this.markUserAsAirdropped()
        }
        return
      }
      return
    }
  }
}
</script>

<style>
  .faucet-div {
    height: 70vh;
    display: grid;
    align-items: center;
    align-content: center;
    justify-content: center;
  }
</style>