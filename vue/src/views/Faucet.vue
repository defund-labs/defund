<template>
  <div>
    <div v-if="store.twitter == null" class="faucet-div" style="height:70vh;" align="center">
      <button style="background-color:#067acc; border-color:#067acc" @click="login" class="sp-button">Verify Twitter</button>
    </div>
    <div v-else class="faucet-div" style="height:60vh;" align="center">
      <h4 style="font-size:x-large; font-weight: bold; margin-bottom: 25px;">You are eligible for the private testnet!</h4>
      <button @click="requestTokens" style="background-color:#6CE5E8; border-color:#6CE5E8" class="sp-button">Request Tokens</button>
    </div>
  </div>
</template>

<script>
import { getAuth, signInWithPopup, TwitterAuthProvider } from "firebase/auth";
import { store } from '../store/local/store.js';
import fs from 'fs';
import Papa from 'papaparse';
export default {
  name: 'Faucet',
  data() {
    const provider = new TwitterAuthProvider();
    const auth = getAuth();
    return {
      auth,
      provider,
      store
    }
  },
  methods:{
    login() {
      signInWithPopup(this.auth, this.provider)
        .then((result) => {
          const credential = TwitterAuthProvider.credentialFromResult(result);
          const user = result.user;
          store.twitter = user
        }).catch((error) => {
          console.error(error)
        });
    },
    checkEligibility() {
      var file = fs.createReadStream("/data/defund_testnet_followers.csv")
      Papa.parse(file, {
          worker: true,
          complete: function(results, file) {
            console.log(results)
            console.log('parsing complete read', count, 'records.'); 
          }
      });
    },
    requestTokens() {
      this.checkEligibility()
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