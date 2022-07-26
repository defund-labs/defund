<template>
    <div class="feedback text">
        <div class="two-grid">
            <div class="check-icon">
            <svg
                width="40"
                height="40"
                viewBox="0 0 64 63"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
            >
                <circle
                cx="32"
                cy="31.5"
                r="29.5"
                stroke="#00CF30"
                stroke-width="4"
                stroke-linecap="round"
                />
                <path
                d="M19 30.1362L28.6557 40L45 23"
                stroke="#00CF30"
                stroke-width="4"
                />
            </svg>
            </div>
            <div @click="clearLastTx" style="font-size: x-large;font-weight: 700;text-align: right;cursor: pointer;">
                x
            </div>
        </div>

        <div style="width: 100%; height: 24px" />

        <div class="tx-feedback-subtitle">
            <slot></slot>
        </div>
    </div>
</template>
 
<script>
import { store } from '../store/local/store.js'
export default {
  name: "Success",
  data() {
    return {
        store: store
    }
  },
  methods: {
    clearLastTx() {
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
</style>