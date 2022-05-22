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
                <h2 class="heading-price">1 ATOM</h2>
                <h4 style="margin-top:10px" class="heading-growth title-spacer">+0.00% (Today)</h4>
            </div>
          </div>
        </div>
        <div></div>
        <div class="w-layout-grid grid-4">
          <div>
            <apexchart height="350" type="area" :options="options" :series="series"></apexchart>
          </div>
          <div>
            <h1 style="margin-top:10px" class="heading-symbol title-spacer">About</h1>
            <p style="font-size: large; font-weight: 400;">{{fund.fund.description}}</p>
          </div>
          <div>
              <h1 style="margin-top:10px; margin-bottom:10px;" class="heading-symbol title-spacer">Composition</h1>
              <v-grid
                theme="material"
                :source="fund.fund.holdings"
                :columns="columns"
                filter=false
                readonly=true
                col-size=275
                row-size=60
                can-focus=false
              ></v-grid>
          </div>
        </div>
      </div>
      <div>
        <div class="buy-div">
          <div style="height: 30px;" class="border-bottom">
            <h4 style="font-size: large; font-weight: 600;">{{"Buy " + fund.fund.symbol}}</h4>
          </div>
          <div style="height: 200px;" class="three-row-grid border-bottom">
            <div style="font-size: large; font-weight: 400;" class="two-col-grid">
              <div>
                Invest In
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
              <div>
                {{computedShares}}
              </div>
            </div>
          </div>
          <div style="display: flex;justify-content: center;">
            <button style="background-color: green; border-color: green;" class="sp-button"> Place Order</button>
          </div>
          <div>

          </div>
        </div>
      </div>
    </div>
    <div v-else>
      Loading...
    </div>
  </div>
</template>

<script>
import { SpSpacer } from '@starport/vue'
import { computed } from 'vue'
import { useStore } from 'vuex'
import { VGrid } from "@revolist/vue3-datagrid";
import _ from 'lodash'
import { store } from '../store/local/popup.js';

export default {
  name: 'Fund',
  components: { 
    SpSpacer,
    VGrid
   },
  data() {
    let $s = useStore()

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
      return currentfundprice
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
      currentPrice: currentFundPriceRaw
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
    }
  },
  methods: {
    updateComputedShares(e) {
      const amount = e.target.value
      this.computedShares = amount / this.currentPrice.amount.amount
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
    height: 400px;
    border-width: 2px;
    border-style: outset;
    border-color: rgba(0, 0, 0, 0.027);
    display: grid;
    grid-template-rows: 0.10fr auto auto;
    padding: 20px;
    background-color: white;
    position: fixed;
    border-radius: 10px;
    z-index: 100;
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
