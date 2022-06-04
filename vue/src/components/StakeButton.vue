<template>
  <div>
    <div class="stake-button-div">
      <button v-on:click="togglePopup()" type="button" class="sp-button stake-button">
        Manage
      </button>  
    </div>
  </div>
</template>
 
<script>
import { store } from '../store/local/store.js';
export default {
  name: "StakeButton",
  data() {
    return {
      stakePopup: store.stakePopup
    }
  },
  methods: {
    togglePopup: function() {
      if (store.stakePopup == false) {
        store.stakePopup = true
      } else {
        store.stakePopup = false
      }
    }
  }
};
</script>

<style>
  .stake-button-div {
    text-align: right;
  }
  .stake-button {
    margin-right: 15px;
    color: white;
    border-style: none;
    background-color: black;
  }
</style>