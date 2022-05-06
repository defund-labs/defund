<template>
    <div>
        <SpSpacer size="sm" />
        <div class="two-col-grid">
          <div style="margin-right:5px;">
              <label :for="`pname`" class="sp-label capitalize-first-letter">
              Name
              </label>
              <input
              :id="`pname`"
              v-model="formData['name']"
              :placeholder="`Enter name`"
              type="text"
              :name="`pname`"
              class="sp-input"
              @change="onFormDataChange"
              />
              <SpSpacer size="xs" />
          </div>
          <div style="margin-left:5px;">
              <label :for="`psymbol`" class="sp-label capitalize-first-letter">
              Symbol
              </label>
              <input
              :id="`psymbol`"
              v-model="formData['symbol']"
              :placeholder="`Enter symbol`"
              type="text"
              :name="`psymbol`"
              class="sp-input"
              @change="onFormDataChange"
              />
              <SpSpacer size="xs" />
          </div>
        </div>
        <div class="two-col-grid">
          <div style="margin-right:5px;">
              <label :for="`pbroker`" class="sp-label capitalize-first-letter">
              Broker Dex
              </label>
              <select
              :id="`pbroker`"
              v-model="formData['broker']"
              :placeholder="`Select broker`"
              type="text"
              :name="`pbroker`"
              class="sp-input"
              @change="onFormDataChange"
              >
                <option value="gdex" selected>Gravity Dex</option>
              </select>
              <SpSpacer size="xs" />
          </div>
          <div style="margin-left:5px;">
              <label :for="`prebalance`" class="sp-label capitalize-first-letter">
              Rebalance Period (Per Block)
              </label>
              <input
              :id="`prebalance`"
              v-model="formData['rebalance']"
              :placeholder="`Enter rebalance period`"
              type="number"
              :name="`prebalance`"
              class="sp-input"
              @change="onFormDataChange"
              />
              <SpSpacer size="xs" />
          </div>
        </div>
        <div :holding="holding.id" v-for="holding in holdings.holdings" v-bind:key="holding.id">
          <div class="three-col-grid">
            <div :holding="holding.id" style="margin-right:5px;">
                  <label :for="`ppoolid`" class="sp-label capitalize-first-letter">
                  Pool
                  </label>
                  <select
                   :holding="holding.id"
                  :id="`ppoolid`"
                  type="number"
                  :name="`ppoolid`"
                  class="sp-input"
                  @change="(el) => updateTokensSelect(el)"
                  >
                    <option v-for="id of getPoolIDs()" v-bind:key="id">
                      {{id}}
                    </option>
                  </select>
            </div>
            <div :holding="holding.id" style="margin-left:5px; margin-right:5px;">
                  <label :for="`ptoken`" class="sp-label capitalize-first-letter">
                  Token
                  </label>
                  <select
                  :holding="holding.id"
                  :id="`ptoken`"
                  type="text"
                  :name="`ptoken`"
                  class="sp-input"
                  @change="(el) => updateTokenOnChange(el)"
                  >
                    <option v-for="token in getTokensFromHolding(holding)" v-bind:key="token">
                      {{token}}
                    </option>
                  </select>
            </div>
            <div :holding="holding.id" style="margin-left:5px;">
                  <label :for="`pholdings`" class="sp-label capitalize-first-letter">
                  % Composition
                  </label>
                  <input
                  :holding="holding.id"
                  :placeholder="`Add % composition`"
                  type="number"
                  class="sp-input"
                  value=10
                  @change="(el) => updateWeightsOnChange(el)"
                  />
            </div>
          </div>
        </div>
        <div style="display: grid;grid-template-columns: auto auto;justify-content: center;">
          <a @click="addHolding" style="font-size: xx-large; font-weight: 600; cursor: pointer; display: flex; justify-content: center; margin-top: 10px;margin-right: 15px;">+</a>
          <a @click="deleteHolding" style="font-size: xx-large; font-weight: 600; cursor: pointer; display: flex; justify-content: center; margin-top: 10px;margin-left: 15px;">-</a>
        </div>
        <SpSpacer size="xs" />
        <div>
              <label :for="`pdescription`" class="sp-label capitalize-first-letter">
              Description
              </label>
              <textarea
              :id="`pdescription`"
              v-model="formData['description']"
              :placeholder="`Enter description`"
              type="text"
              :name="`pdescription`"
              class="sp-input"
              @change="onFormDataChange"
              />
        </div>
        <SpSpacer size="xs" />
        <div style="display: grid;grid-template-columns: auto auto;">
          <SpButton
            type="primary"
            @click="submitCreateFund"
            style="margin-right: 10px;"
            :disabled="createButtonDisabled"
          >
            Create Fund
          </SpButton>
          <SpButton
            type="secondary"
            style="margin-left: 10px;"
            @click="router.push('/funds')"
          >
            Cancel
          </SpButton>
        </div>
    </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive } from 'vue'
import { useStore } from 'vuex'
import { store } from '../store/local/popup'
import router from '../router'
import { SpButton, SpDropdown, SpSpacer, SpTypography } from '@starport/vue'
import _ from 'lodash';

export default defineComponent({
  name: 'CreateFund',

  components: {
    SpSpacer,
    SpTypography,
    SpButton,
    SpDropdown,
  },
  data(props) {
    let $s = useStore()
    let creator = computed(() => $s.getters['common/wallet/address'])
    let formData = reactive({
      broker: "gdex",
      name: "",
      symbol: "",
      rebalance: 100,
      description: ""
    })

    let holdings = reactive({
      holdings: [
        {
          "id": 1,
          "token": "OSMO",
          "pool": "1",
          "weight": "10"
        }
      ]
    })

    const pools = JSON.parse(JSON.stringify(store.pools))

    return {
      creator: creator,
      formData,
      holdings: holdings,
      pools: pools,
      store: $s,
      createButtonDisabled: true,
      router,
      local: store
    }
  },
  methods: {
    addHolding() {
      const id = this.holdings.holdings[this.holdings.holdings.length - 1].id + 1
      this.holdings.holdings.push({
          "id": id,
          "token": "OSMO",
          "pool": "1",
          "weight": "10"
        })
    },
    editHolding() {
    },
    deleteHolding() {
      if (this.holdings.holdings.length > 1) {
        this.holdings.holdings.pop()
      }
    },
    getPoolIDs() {
      var ids = []
      for (const pool of this.pools) {
        ids.push(pool.id)
      }
      return ids
    },
    getTokensFromHolding(holding) {
      const pool = _.filter(this.pools, (p) => {
        return p.id == holding.pool
      })
      var denomsRaw
      var denoms = []
      if (pool.length > 0) { 
        denomsRaw = pool[0].reserve_coin_denoms
        for (const denom of denomsRaw) {
          denoms.push(store.IBCToTokenMap[denom])
        }
        return denoms
      }
      return null
    },
    async submitCreateFund() {
      var fullHoldingString = ""
      var connectionId
      var baseDenom
      for (const holding of this.holdings.holdings) {
        const stringHolding = store.TokenToIBCMap[holding.token] + ":" + holding.weight + ":" + holding.pool
        if (fullHoldingString == "") {
          fullHoldingString = stringHolding
        } else {
          fullHoldingString = fullHoldingString + "," + stringHolding
        }
      }

      // Set connection, baseDenom to gdex params based on broker (if gdex)
      if (this.formData.broker == "gdex") {
        connectionId = "connection-0"
        baseDenom = "uatom"
      }

      const value = {
        creator: this.creator,
        symbol: this.formData.symbol,
        name: this.formData.name,
        description: this.formData.description,
        broker: this.formData.broker,
        holdings: fullHoldingString,
        rebalance: this.formData.rebalance,
        baseDenom: baseDenom,
        connectionId: connectionId
      }

      const res = await this.store.dispatch("defundlabs.defund.etf/sendMsgCreateFund", {
          value: value,
          fee: [{
            amount: "200000",
            denom: "ufetf"
          }],
          memo: ""
      })

      console.log(res)

      if(res.code == 0) { 
        this.local.showTxSuccess = true 
        this.local.showTxFail = false
        this.local.lastTxHash = res.transactionHash
        this.local.showTxStatus = true
      } else {
        this.local.showTxSuccess = false
        this.local.showTxFail = true
        this.local.lastTxHash = res.transactionHash
        this.local.lastTxLog = res.rawLog
        this.local.showTxStatus = true
      }

      return res
    },
    updateTokensSelect(el) {
      const index = el.target.getAttribute("holding") - 1
      var holding = this.holdings.holdings[index]
      holding.pool = el.target.value
    },    
    updateWeightsOnChange(el) {
      const index = el.target.getAttribute("holding") - 1
      var holding = this.holdings.holdings[index]
      holding.weight = el.target.value
    },
    updateTokenOnChange(el) {
      const index = el.target.getAttribute("holding") - 1
      var holding = this.holdings.holdings[index]
      holding.token = el.target.value
    },
    onFormDataChange() {
      const creator = this.creator != null && this.creator != "" 
      const symbol = this.formData.symbol != null && this.formData.symbol != "" 
      const name = this.formData.name != null && this.formData.name != ""
      const description = this.formData.description != null && this.formData.description != ""
      const broker = this.formData.broker != null && this.formData.broker != ""
      const rebalance = this.formData.rebalance != null && this.formData.rebalance != ""

      this.createButtonDisabled = !(creator && symbol && name && description && broker && rebalance)
    }
  }
})
</script>

<style scoped lang="scss">

.sp-label {
  display: block;
  text-align: left;
  width: 100%;
  margin: 0 4px;

  font-family: Inter;
  font-style: normal;
  font-weight: normal;
  font-size: 13px;
  line-height: 153.8%;
  /* identical to box height, or 20px */

  /* light/muted */

  color: rgba(0, 0, 0, 0.667);
}
.sp-input {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 12px 16px;
  width: 100%;
  height: 48px;
  background: rgba(0, 0, 0, 0.03);
  border-radius: 10px;
  margin: 4px 0px;
  border: 0;
}

.capitalize-first-letter:first-letter {
  text-transform: uppercase;
}
.two-col-grid {
  display: grid;
  grid-template-columns: 50% 50%;
}
.three-col-grid {
  display: grid;
  grid-template-columns: 33.34% 33.33% 33.33%;
}
#pdescription{
  height: 150px;
}
</style>
