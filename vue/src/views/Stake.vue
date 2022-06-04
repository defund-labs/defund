<template>
  <div class="child-container">
    <div v-if="store.stakePopup">
      <StakePopup/>
    </div>
    <div id="your-stake-div">
      <header class="funds-header">
        <h2 class="title">Your Stake</h2>
      </header>
      <div class="grid-div">
        <v-grid
          style="min-height: 300px;"
          theme="material"
          :source="rows_rewards"
          :columns="columns_rewards"
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
    <div style="margin-top: 25px;" id="validators-div">
      <header class="funds-header">
        <h2 class="title">All Validators</h2>
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
        <div class="two-col-grid">
          <div align="left">
            <button :disabled="page == 0" @click="backPage" class="sp-button">Back</button>
          </div>
          <div align="right">
            <button :disabled="page == (Math.ceil(Number(total)/100) - 1)" @click="nextPage" class="sp-button">Next</button>
          </div>
        </div>
      </div>
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
import { store } from '../store/local/store.js';

export default {
  name: "Stake",
  data() {
    let $s = useStore()

    let address = computed(() => {
        return $s.getters['common/wallet/address']
    })

    let rewards = computed(() => {
      $s.dispatch("cosmos.staking.v1beta1/QueryDelegatorValidators", {params: { delegator_addr: address.value }, subscribe: false, all: false })
      var rewards_raw = $s["getters"]["cosmos.staking.v1beta1/getDelegatorValidators"]({params: { delegator_addr: address.value }})
      var rewards = rewards_raw["validators"]
      if (typeof(rewards) != "undefined") {
        for (let i = 0; i < rewards.length; i++) {
          /////////////// Add Rewards /////////////////
          $s.dispatch("cosmos.distribution.v1beta1/QueryDelegationRewards", {params: { delegator_address: address.value, validator_address: rewards[i].operator_address }, subscribe: false, all: false })
          var rewards_for_val_raw = $s["getters"]["cosmos.distribution.v1beta1/getDelegationRewards"]({params: { delegator_address: address.value, validator_address: rewards[i].operator_address }})
          rewards_for_val_raw = _.filter(rewards_for_val_raw["rewards"], function(o) { return o.denom == "ufetf" })
          rewards[i]["rewards"] = rewards_for_val_raw
          /////////////////////////////////////////////
          rewards[i]["tokens"] = String(Number(rewards[i]["tokens"])/1000000) + " FETF"
          rewards[i] = flatten(rewards[i]);
          rewards[i]["commission.commission_rates.rate"] = String(_.round(Number(rewards[i]["commission.commission_rates.rate"]) * 100, 2)) + "%"
          rewards[i]["delegator_shares"] = String(_.round(Number(rewards[i]["delegator_shares"])/1000000, 2)) + " FETF"
          rewards[i]["rewards.0.amount"] = String(_.round(Number(rewards[i]["rewards.0.amount"])/1000000, 2)) + " FETF"
          if (rewards[i]["rewards.0.amount"] == "NaN FETF") { rewards[i]["rewards.0.amount"] = "0" + " FETF" }
        }
      }
      return rewards
    })

    let vals = computed(() => {
      $s.dispatch("cosmos.staking.v1beta1/QueryValidators", {subscribe: false, all: false})
      var validators_raw = $s["getters"]["cosmos.staking.v1beta1/getValidators"]()
      var validators = validators_raw["validators"]
      validators = _.orderBy(validators, [function(o) { return Number(o.tokens); }], ["desc"]);
      if (typeof(validators) != "undefined") {
        for (let i = 0; i < validators.length; i++) {
          validators[i]["tokens"] = String(_.round(Number(validators[i]["tokens"])/1000000, 2)) + " FETF"
          validators[i] = flatten(validators[i]);
          validators[i]["commission.commission_rates.rate"] = String(_.round(Number(validators[i]["commission.commission_rates.rate"]) * 100, 2)) + "%"
        }
      }
      if(typeof(validators_raw) == "undefined") { this.total = Number(validators_raw["pagination"]["total"]) }
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
              "margin-right": "32px",
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
      columns_rewards: [
        { name: "Moniker", prop: "description.moniker"},
        { name: "Website", prop: "description.website"},
        { name: "Delegations", prop: "delegator_shares"},
        { name: "Rewards", prop: "rewards.0.amount"},
        { cellTemplate: (createElement, props) => { return createElement('div', {
          style: {
            "text-align": "right",
            "margin-right": "25px"
          }
        }, createElement('button', {
          class: "sp-button sp-button-button", style: {
              "margin-right": "10px",
              "color": "white",
              "border-style": "none",
              "background-color": "black"
           }, onClick() {
                if (store.stakePopup == false) {
                store.stakePopup = true
                $s.dispatch("cosmos.staking.v1beta1/QueryValidators", {subscribe: false, all: false})
                var validators_raw = $s["getters"]["cosmos.staking.v1beta1/getValidators"]()
                store.currentValidator = props.model.validator_address
              } else {
                store.stakePopup = false
              }
           }
        }, "Claim"),
        createElement('button', {
          class: "sp-button sp-button-button", style: {
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
        }, "Manage"))
        }},
      ],
      rows_rewards: rewards,
      rows: vals,
      store: store,
      s: $s,
      page: 0,
      total: null
    };
  },
  methods: {
    async nextPage() {
      let $s = useStore()

      this.page = this.page + 1
      var validators_raw = await $s.dispatch("cosmos.staking.v1beta1/QueryValidators", {query: { "pagination.offset": this.page * 100 }, subscribe: false, all: false})
      var validators = await validators_raw["validators"]
      validators = await _.orderBy(validators, [function(o) { return Number(o.tokens); }], ["desc"]);
      if (typeof(validators) != "undefined") {
        for (let i = 0; i < validators.length; i++) {
          validators[i]["tokens"] = String(Number(validators[i]["tokens"])/1000000) + " FETF"
          validators[i] = flatten(validators[i]);
          validators[i]["commission.commission_rates.rate"] = String(_.round(Number(validators[i]["commission.commission_rates.rate"]) * 100, 2)) + "%"
        }
      }
      this.rows = computed(() => { return JSON.parse(JSON.stringify(validators)) })
    },
    async backPage () {
      let $s = useStore()

      this.page = this.page - 1
      var validators_raw = await $s.dispatch("cosmos.staking.v1beta1/QueryValidators", {query: { "pagination.offset": this.page * 100 }, subscribe: false, all: false})
      var validators = await validators_raw["validators"]
      validators = await _.orderBy(validators, [function(o) { return Number(o.tokens); }], ["desc"]);
      if (typeof(validators) != "undefined") {
        for (let i = 0; i < validators.length; i++) {
          validators[i]["tokens"] = String(Number(validators[i]["tokens"])/1000000) + " FETF"
          validators[i] = flatten(validators[i]);
          validators[i]["commission.commission_rates.rate"] = String(_.round(Number(validators[i]["commission.commission_rates.rate"]) * 100, 2)) + "%"
        }
      }
      this.rows = computed(() => { return JSON.parse(JSON.stringify(validators)) })
    }
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
  .child-container {
    padding-right: 32px;
    padding-left: 32px;
    padding-bottom: 32px;
  }
</style>

