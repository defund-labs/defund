<template>
  <div class="container">
    <div class="w-layout-grid grid">
      <div class="w-layout-grid grid-3">
        <div>
          <div class="w-layout-grid grid-fund-header">
            <div>
                <h1 class="heading-title">{{fund.name}}</h1>
                <h1 style="margin-top:10px" class="heading-symbol title-spacer">{{fund.symbol}}</h1>
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
            <p>{{fund.description}}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { SpSpacer } from '@starport/vue'
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'Fund',
  components: { 
    SpSpacer
   },
  data: function() {
    let $s = useStore()

    const symbol = this.$route.params.symbol
    let fundRaw = computed(() => {
      $s.dispatch("defundlabs.defund.etf/QueryFund", { params: { symbol: symbol }, subscribe: false, all: false })
      var fund = $s["getters"]["defundlabs.defund.etf/getFund"]({
        params: { symbol: symbol }
      })["fund"]
      return fund
    })
    const fund = JSON.parse(JSON.stringify(fundRaw.value))

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
          categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998]
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
        dataLabels: {
          enabled: false
        }
      },
      series: [{
        name: 'Price',
        data: [30, 40, 45, 50, 49, 60, 70, 91]
      }],
      fund: fund
    }
  }
}
</script>

<style>
  .w-layout-grid {
    display: -ms-grid;
    display: grid;
    grid-auto-columns: 1fr;
    -ms-grid-columns: 1fr 1fr;
    grid-template-columns: 1fr 1fr;
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
</style>
