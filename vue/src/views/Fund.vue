<template>
  <div class="container">
    <div v-if="fund.fund" class="w-layout-grid">
      <div class="w-layout-grid grid-3">
        <div>
          <div class="w-layout-grid grid-fund-header">
            <div>
                <h1 class="heading-title">{{fund.fund.name}}</h1>
                <h1 style="margin-top:10px" class="heading-symbol title-spacer">{{fund.fund.symbol}}</h1>
            </div>
            <div>
                <h2 class="heading-price">{{currentPrice}} {{store.IBCToTokenMap[fund.fund.baseDenom]}}</h2>
                <h4 style="margin-top:10px" class="heading-growth title-spacer">+0.00% (Today)</h4>
            </div>
          </div>
        </div>
        <div></div>
        <div class="w-layout-grid grid-4">
          <div v-if="series[0].data.length > 0">
            <apexchart height="350" type="area" :options="options" :series="series"></apexchart>
          </div>
          <div>
            <h1 style="margin-top:10px" class="heading-symbol title-spacer">About</h1>
            <p style="font-size: large; font-weight: 400;">{{fund.fund.description}}</p>
          </div>
          <div>
              <h1 style="margin-top:10px; margin-bottom:10px;" class="heading-symbol title-spacer">Composition</h1>
              <v-grid
                style="height: fit-content;"
                theme="material"
                :source="holdings"
                :columns="columns"
                filter=false
                readonly=true
                col-size=275
                row-size=60
                can-focus=false
                height="fit-content"
              ></v-grid>
          </div>
        </div>
      </div>
      <div>
        <div v-if="buyDiv" class="buy-div">
          <div style="height: 100%; margin-bottom: 15px;" class="two-col-grid border-bottom">
            <h4 style="font-size: large; font-weight: 600;">Create</h4>
            <div style="text-align: right;">
              <img @click="toggleBuySell" style="text-align: right; grid-template-columns: 0.80fr 0.20fr; cursor: pointer" height="15" src="/redo.svg"/>
            </div>
          </div>
          <div style="height: 200px;" class="three-row-grid">
            <div style="font-size: large; font-weight: 400;" class="two-col-grid">
              <div>
                Base
              </div>
              <div>
                <select v-model="selected" style="width: 100%; height: 36px; border-radius: 5px; padding: 0px 10px;" class="input input-border">
                  <option :value="fund.fund.baseDenom">{{mapToIBC[fund.fund.baseDenom]}}</option>
                </select>
              </div>
            </div>
            <div style="font-size: large; font-weight: 400;" class="two-col-grid">
              <div>
                Amount
              </div>
              <div>
                <input @change="(event) => updateComputedShares(event)" type="number" style="width: 100%; height: 36px; border-radius: 5px; padding: 0px 10px;" class="input input-border"/>
              </div>
            </div>
            <div style="font-size: large; font-weight: 400;" class="two-col-grid">
              <div>
                Shares
              </div>
              <div style="text-align: right;">
                {{computedShares}}
              </div>
            </div>
          </div>
          <div class="border-bottom" style="display: flex;justify-content: center;padding-bottom: 15px;">
            <button @click="createShareAction" style="background-color: green; border-color: green;" class="sp-button">Create Shares</button>
          </div>
          <div style="font-size: medium; font-weight: 400; margin-top: 10px;" class="two-col-grid">
            <div>
              Available
            </div>
            <div style="text-align: right;">
              {{getDenomBalance(balances, fund.fund.baseDenom)}}
            </div>
          </div>
        </div>
        <div v-else class="sell-div">
          <div style="height: 100%; margin-bottom: 15px;" class="two-col-grid border-bottom">
            <h4 style="font-size: large; font-weight: 600;">Redeem</h4>
            <div style="text-align: right;">
              <img @click="toggleBuySell" style="text-align: right; grid-template-columns: 0.80fr 0.20fr; cursor: pointer;" height="15" src="/redo.svg"/>
            </div>
          </div>
          <div style="height: 300px;" class="three-row-grid">
            <div style="font-size: large; font-weight: 400;" class="two-col-grid">
              <div>
                Receive
              </div>
              <div>
                <select v-model="selected" style="width: 100%; height: 36px; border-radius: 5px; padding: 0px 10px;" class="input input-border">
                  <option :value="fund.fund.baseDenom">Shares</option>
                </select>
              </div>
            </div>
            <div style="font-size: large; font-weight: 400;" class="two-col-grid">
              <div>
                Amount
              </div>
              <div>
                <input @change="(event) => updateComputedSharesSell(event)" type="number" style="width: 100%; height: 36px; border-radius: 5px; padding: 0px 10px;" class="input input-border"/>
              </div>
            </div>
            <div class="border-bottom" style="display: flex;justify-content: center;">
              <button @click="redeemShareAction" style="background-color: red; border-color: red;" class="sp-button">Redeem Shares</button>
            </div>
            <div style="font-size: medium; font-weight: 400; margin-top: 10px;" class="two-col-grid">
              <div>
                Available
              </div>
              <div style="text-align: right;">
                {{getFundDenomBalance(balances, fund.fund.shares.denom)}}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div style="display:grid; justify-content: center; align-items: center; font-size: medium; height: 70vh; font-weight: 600;" v-else>
      <img
          width="50"
          src="/logo_animation.gif"
      />
    </div>
  </div>
</template>

<script>
import { SpSpacer } from '@starport/vue'
import { computed } from 'vue'
import { useStore } from 'vuex'
import { VGrid } from "@revolist/vue3-datagrid";
import _ from 'lodash'
import { store } from '../store/local/store.js';

export default {
  name: 'Fund',
  components: { 
    SpSpacer,
    VGrid
   },
  data() {
    let $s = useStore()

    let address = computed(() => {
        return $s.getters['common/wallet/address']
    })

    // Add fund data to store
    const symbol = this.$route.params.symbol
    let fundRaw = computed(() => {
      $s.dispatch("defundlabs.defund.etf/QueryFund", { params: { symbol: symbol }, subscribe: false, all: false })
      var fund = $s["getters"]["defundlabs.defund.etf/getFund"]({
        params: { symbol: symbol }
      })
      return fund
    })

    // Add fund price data to store
    let fundPriceRaw = computed(() => {
      $s.dispatch("defundlabs.defund.etf/QueryFundPriceAll", { params: { symbol: symbol }, query: { "pagination.reverse": true, "pagination.limit": 100 }, subscribe: false, all: false })
      var fundprice = $s["getters"]["defundlabs.defund.etf/getFundPriceAll"]({
        params: { symbol: symbol },
        query: { "pagination.reverse": true, "pagination.limit": 100 }
      })
      return fundprice
    })

    // Add current fund price data to store
    let currentFundPriceRaw = computed(() => {
      $s.dispatch("defundlabs.defund.etf/QueryFundPrice", { query: { symbol: symbol }, subscribe: true, all: false })
      var currentfundprice = $s["getters"]["defundlabs.defund.etf/getFundPrice"]({
        params: { symbol: symbol }
      })
      return currentfundprice.value ? currentfundprice : 1
    })

    // Add wallet balances to store
    let balances = computed(() => {
        if (address.value) {
            $s.dispatch('cosmos.bank.v1beta1/QueryAllBalances', {
                params: { address: address.value },
                options: { subscribe: true }
            })
        }

        return $s.getters['cosmos.bank.v1beta1/getAllBalances']({
            params: { address: address.value },
        })
    })

    return {
      options: {
        chart: {
          id: 'fund-price-historical',
          animations: {
            enabled: false
          },
          toolbar: {
            show: false
          },
          fontFamily: "inherit",
          foreColor: "inherit"
        },
        xaxis: {
          labels: {
            show: false
          },
          axisBorder: {
            show: false
          },
          axisTicks: {
            show: false
          },
          categories: [],
        },
        yaxis: {
          labels: {
            show: false
          },
          axisBorder: {
            show: false
          },
          axisTicks: {
            show: false
          }
        },
        grid: {
            show: false
        },
        stroke: {
          curve: "smooth"
        },
        tooltip: {
          followCursor: true
        },
        dataLabels: {
          enabled: false
        }
      },
      series: [],
      fund: fundRaw,
      fundprice: fundPriceRaw,
      columns: [
        { name: "Token", prop: "token" },
        { name: "Percent (%)", prop: "percent" }, 
      ],
      mapToIBC: store.IBCToTokenMap,
      selected: 'uatom',
      computedShares: 0,
      computedSharesSell: 0,
      currentPrice: currentFundPriceRaw,
      address,
      balances,
      store,
      buyDiv: true,
      holdings: []
    }
  },
  watch: {
    fundprice(newValue, oldValue) {
      if(this.series.length <= 0) {
        this.options.xaxis.categories = _.map(this.fundprice.price, (p) => { 
          return Number(p.height)
        }).sort()
        this.series.push({
          name: 'Price',
          data: _.map(this.fundprice.price, (p) => { 
            return Number(p.amount.amount)/1000000 
          })
        })
      }
    },
    fund(newValue, oldValue) {
      var holdings = newValue.fund.holdings
      for (let i = 0; i < holdings.length; i++) {
        holdings[i]["token"] = this.mapToIBC[holdings[i]["token"]]
      }
      this.holdings = holdings
      return holdings
    }
  },
  methods: {
    updateComputedShares(e) {
      const amount = e.target.value
      this.computedShares = amount / this.currentPrice
    },
    updateComputedSharesSell(e) {
      const amount = e.target.value
      this.computedSharesSell = amount * this.currentPrice
    },
    getDenomBalance(balances, denom) {
      var balance = null

      if (balances) {

          balances = JSON.parse(JSON.stringify(balances)).balances

          const found = _.filter(balances, function(o) { return o.denom == denom })

          if(found.length == 0) { return String(0) + " " + this.store.IBCToTokenMap[denom] }

          balance = String(Number(found[0].amount)/1000000)

          return balance + " " + this.store.IBCToTokenMap[denom]

      }

      return String(0)
    },
    getFundDenomBalance(balances, denom) {
      var balance = null

      if (balances) {

          balances = JSON.parse(JSON.stringify(balances)).balances

          const found = _.filter(balances, function(o) { return o.denom == denom })

          if(found.length == 0) { return String(0) + " Shares" }

          balance = String(Number(found[0].amount)/1000000)

          return balance + " " + "Shares"

      }

      return String(0)
    },
    toggleBuySell() {
      if (this.buyDiv == false) {
        this.buyDiv = true
      } else {
        this.buyDiv = false
      }
    },
    redeemShareAction() {
      alert("Redemptions will be activated on the first chain upgrade...")
    },
    createShareAction() {
      alert("Creations will be activated on the first chain upgrade...")
    }
  },
}
</script>

<style>
  .w-layout-grid {
    display: -ms-grid;
    display: grid;
    grid-auto-columns: 1fr;
    grid-template-columns: 0.75fr;
    -ms-grid-rows: auto auto;
    grid-template-rows: auto auto;
    grid-row-gap: 16px;
    grid-column-gap: 16px;
  }

  .grid {
    -ms-grid-columns: 1fr 0.5fr;
    grid-template-columns: 1fr 0.5fr;
    -ms-grid-rows: auto;
    grid-template-rows: auto;
  }

  .grid-fund-header {
    -ms-grid-columns: 1fr 1fr;
    grid-template-columns: 1fr 1fr;
    -ms-grid-rows: auto;
    grid-template-rows: auto;
  }

  .grid-3 {
    -ms-grid-columns: 1fr;
    grid-template-columns: 1fr;
    -ms-grid-rows: auto auto auto;
    grid-template-rows: auto auto auto;
  }

  .heading-title {
    font-size: 2.50rem;
    font-weight: 730;
  }

  .heading-symbol {
    font-size: 2.00rem;
    font-weight: 600;
  }

  .heading-price {
    font-size: 2.25rem;
    font-weight: 650;
    text-align: right;
  }

  .heading-growth {
    font-size: 1.75rem;
    font-weight: 500;
    color: green;
    text-align: right;
  }

  .grid-4 {
    -ms-grid-columns: 1fr;
    grid-template-columns: 1fr;
  }

  h2 {
    margin: 0px;
    line-height: 1.15;
  }

  h4 {
    margin: 0px;
    line-height: 1rem;
  }

  .chart {
      --chart-positive: var(--positive);
      --chart-negative: var(--negative);
  }

  revo-grid {
      min-height: 200px;
  }

  .buy-div {
    right: 30px;
    top: 20%;
    width: 22%;
    border-width: 2px;
    border-style: outset;
    border-color: rgba(0, 0, 0, 0.027);
    display: grid;
    grid-template-rows: 0.10fr auto auto;
    padding: 20px;
    background-color: white;
    position:fixed;
    border-radius: 10px;
    z-index: 80;
  }
  .sell-div {
    right: 30px;
    top: 20%;
    width: 22%;
    border-width: 2px;
    border-style: outset;
    border-color: rgba(0, 0, 0, 0.027);
    display: grid;
    grid-template-rows: 0.10fr auto auto;
    padding: 20px;
    background-color: white;
    position: fixed;
    border-radius: 10px;
    z-index: 80;
  }
  .border-bottom {
    border-color: rgba(0, 0, 0, 0.027);
    border-width: 2px;
    border-bottom-style: solid;
  }
  .three-row-grid {
    display: grid;
    grid-template-rows: auto auto auto;
  }
  .two-col-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    align-items: center;
  }
  .input-border {
    border-color: rgba(0, 0, 0, 0.027);
    border-width: 2px;
    border-style: solid;
  }
</style>
