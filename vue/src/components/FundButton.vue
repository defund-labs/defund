<template>
  <div class="fund-button-div">
      <a @click="router.push(`/funds/${symbol}`)" type="button" class="sp-button fund-button">
        -->
      </a>
  </div>
</template>
 
<script>
import router from '../router'
export default {
  name: "FundButton",
  props: ["rowIndex", "prop", "data"],
  data() {
      return {
          symbol: this.data[this.rowIndex][this.prop],
          router: router
      }
  }
};
</script>

<style>
  .fund-button-div {
    text-align: right;
  }
  .fund-button {
    margin-right: 15px;
    color: white;
    border-style: none;
    background-color: black;
  }
</style>