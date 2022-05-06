<template>
  <div class="container">
    <header class="funds-header">
      <h2 class="title">Funds</h2>
      <a @click="router.push('/funds/create')" style="font-size: large;font-weight: 600;cursor: pointer;">Create -></a>
    </header>
    <v-grid
      theme="material"
      :source="rows"
      :columns="columns"
      filter=true
      readonly=true
      resize=true
      col-size=165
      row-size=60
      can-focus=false
    ></v-grid>
  </div>
</template>
 
<script>
import { VGrid, VGridVueTemplate } from "@revolist/vue3-datagrid";
import { computed } from 'vue';
import { useStore } from 'vuex';
import { SpTheme, SpButton } from '@starport/vue';
import FundButton from '../components/FundButton.vue'
import router from '../router'

export default {
  name: "Funds",
  data() {
    let $s = useStore()

    let funds = computed(() => {
      $s.dispatch("defundlabs.defund.etf/QueryFundAll", {subscribe: false, all: false})
      var allFunds = $s["getters"]["defundlabs.defund.etf/getFundAll"]()["fund"]
      return allFunds
    })

    return {
      columns: [
        { name: "Name", prop: "name", sortable: true },
        { name: "Symbol", prop: "symbol", sortable: true }, 
        { name: "Broker", prop: "broker", sortable: true }, 
        { name: "Price", cellTemplate: (createElement, props) => { return createElement('span', {}, "1 ATOM" )} },
        { name: "Market Cap", cellTemplate: (createElement, props) => { return createElement('span', {}, "0 ATOM" )} },
        { name: "1 Day", cellTemplate: (createElement, props) => { return createElement('span', { style: { color: "green" } }, "0%")} },
        { prop: "symbol", cellTemplate: VGridVueTemplate(FundButton) },
      ],
      rows: funds,
      router: router
    };
  },
  components: {
    VGrid, SpTheme, SpButton, FundButton
  },
};
</script>

<style>
  .funds-header {
    display: flex;
    align-content: center;
    align-items: center;
    justify-content: space-between;
  }
  revo-grid {
    min-height: 500px;
  }
  .title {
    font-family: Inter, serif;
    font-style: normal;
    font-weight: 600;
    font-size: 28px;
    line-height: 127%;
    /* identical to box height, or 36px */

    letter-spacing: -0.02em;
    font-feature-settings: 'zero';

    color: #000000;
    margin-top: 0;
  }
  revo-grid[theme=material] {
      font-family: "Inter", sans-serif;
  }
  revo-grid[theme=material] revogr-data .rgRow {
    line-height: 60px;
  }
  revo-grid[theme=material] revogr-data .rgRow:hover {
      background-color: rgba(233, 234, 237, 0.5);
      cursor: pointer
  }
</style>
