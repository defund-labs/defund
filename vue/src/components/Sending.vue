<template>
    <div class="feedback text">
        <div class="two-grid">
            <div class="check-icon">
            <img
                width="50"
                src="/logo_animation.gif"
            />
            </div>
            <div @click="clearLastTx" style="font-size: x-large;font-weight: 700;text-align: right;cursor: pointer;">
                x
            </div>
        </div>

        <div style="width: 100%; height: 24px" />

        <div class="tx-feedback-subtitle">
            Submitting Transaction...
        </div>
    </div>
</template>
 
<script>
import { store } from '../store/local/store.js'
export default {
  name: "Sending",
  data() {
    return {
        store: store
    }
  },
  methods: {
    clearLastTx() {
        this.store.sendingTx = false
        this.store.lastTxHash = null
        this.store.lastTxLog = null
        this.store.showTxSuccess = false
        this.store.showTxFail = false
        this.store.showTxStatus = false
    }
  }
};
</script>

<style>
    .text {
        font-size: large;
    }
    .two-grid {
        display: grid;
        grid-template-columns: auto auto;
    }
    .tx-feedback-subtitle {
        font-weight: 600;
        font-size: medium;
    }
</style>