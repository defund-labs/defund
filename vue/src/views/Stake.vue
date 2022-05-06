<template>
  <div class="container">
    <div v-if="store.stakePopup">
      <StakePopup/>
    </div>
    <header class="funds-header">
      <h2 class="title">Validators</h2>
    </header>
    <div class="grid-div">
      <v-grid
        theme="material"
        :source="rows"
        :columns="columns"
        filter=true
        readonly=true
        resize=true
        col-size=200
        row-size=60
        can-focus=false
        row-headers=true
      ></v-grid>
    </div>
  </div>
</template>
 
<script>
import VGrid, {VGridVueTemplate} from "@revolist/vue3-datagrid";
import StakeButton from '../components/StakeButton.vue'
import StakePopup from '../components/StakePopup.vue'
import { computed } from 'vue';
import { useStore } from 'vuex';
import { SpTheme } from '@starport/vue';
import flatten from 'flat';
import _ from 'lodash';
import { store } from '../store/local/popup.js';

export default {
  name: "Stake",
  data() {
    let $s = useStore()

    let vals = computed(() => {
    $s.dispatch("cosmos.staking.v1beta1/QueryValidators", {subscribe: false, all: false})
    var validators = $s["getters"]["cosmos.staking.v1beta1/getValidators"]()["validators"]
    validators = _.sortBy(validators, [function(o) { return Number(o.tokens); }]);
    if (typeof(validators) != "undefined") {
      for (let i = 0; i < validators.length; i++) {
        validators[i]["tokens"] = String(Number(validators[i]["tokens"])/1000000) + " FETF"
        validators[i] = flatten(validators[i]);
        validators[i]["commission.commission_rates.rate"] = String(_.round(Number(validators[i]["commission.commission_rates.rate"]) * 100, 2)) + "%"
      }
    }
    return validators
    })

    return {
      columns: [
        { name: "Moniker", prop: "description.moniker"},
        { name: "Website", prop: "description.website"},
        { name: "Staked", prop: "tokens"},
        { name: "Commission", prop: "commission.commission_rates.rate" },
        { cellTemplate: (createElement, props) => { return createElement('div', {
          style: {
            "text-align": "right"
          }
        }, createElement('button', {
          class: "sp-button sp-button-button", style: {
              "margin-right": "15px",
              "color": "white",
              "border-style": "none",
              "background-color": "black"
           }, onClick() {
                if (store.stakePopup == false) {
                store.stakePopup = true
                store.currentValidator = props.model
              } else {
                store.stakePopup = false
              }
           }
        }, "Manage"))}}
        ],
      rows: vals,
      store: store
    };
  },
  components: {
    VGrid, SpTheme, StakeButton, StakePopup
  }
};
</script>

<style>
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

