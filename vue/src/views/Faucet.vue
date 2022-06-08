<template>
  <div>
    <div v-if="store.twitter == null" class="faucet-div" style="height:70vh;" align="center">
      <button style="background-color:#067acc; border-color:#067acc" @click="login" class="sp-button">Verify Twitter</button>
    </div>
    <div v-else class="faucet-div" style="height:60vh;" align="center">
      <h4 v-if="eligible" style="font-size:x-large; font-weight: bold; margin-bottom: 25px;">You are eligible for the private testnet!</h4>
      <h4 v-if="!eligible" style="font-size:x-large; font-weight: bold; margin-bottom: 25px;">You are not eligible for the private testnet.</h4>
      <button v-if="eligible" :disabled="!eligible ||	airdropped" @click="requestTokens" style="background-color:#6CE5E8; border-color:#6CE5E8" class="sp-button">{{airdropped ? "Already Airdropped" : "Request Tokens"}}</button>
    </div>
  </div>
</template>

<script>
import { getAuth, signInWithPopup, TwitterAuthProvider } from "firebase/auth";
import { store } from '../store/local/store.js';
import { computed, onRenderTriggered } from 'vue';
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

    const provider = new TwitterAuthProvider();
    const auth = getAuth();
    let address = computed(() => {
        return $s.getters['common/wallet/address']
    })

    return {
      auth,
      provider,
      store,
      eligible: false,
      airdropped: false,
      address,
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
      const response = await this.axios.get("http://0u8e8hurclf2935mlnfl1t770g.akashi.derateknoloji.com/users/" + store.twitter.providerData[0].uid)
      this.eligible = await response.data.eligible
      this.airdropped = await response.data.data.airdropped
      return response
    },
    async markUserAsAirdropped() {
      const response = await this.axios.put("http://0u8e8hurclf2935mlnfl1t770g.akashi.derateknoloji.com/users", { "id": store.twitter.providerData[0].uid })
      return response
    },
    async requestTokens() {
      const res = await this.checkEligibility()
      const eligible = await res.data.eligible
      if(eligible) {
        const ret = await this.axios.post('http://5k74k6pn1lftt9giu8hj6r638k.ingress.bdl.computer/http://5t2e9or7jdcfpelef01sji8dj4.ingress.akash.virtualhosting.hk', {
          address: this.address,
          coins: [
            "100000000ufetf"
          ]
        })
        console.log(ret)
        this.markUserAsAirdropped()
        return ret
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