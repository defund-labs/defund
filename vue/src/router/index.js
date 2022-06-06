import { createRouter, createWebHashHistory } from 'vue-router'

import Portfolio from '../views/Portfolio.vue'
import Funds from '../views/Funds.vue'
import Fund from '../views/Fund.vue'
import Stake from '../views/Stake.vue'
import Gov from '../views/Gov.vue'
import Faucet from '../views/Faucet.vue'
import CreateFunds from '../views/CreateFunds.vue'

const routerHistory = createWebHashHistory()
const routes = [
  { path: '/', component: Portfolio },
  { path: '/portfolio', component: Portfolio },
  { path: '/funds', component: Funds },
  { path: '/funds/:symbol', component: Fund },
  { path: '/funds/create', component: CreateFunds },
  { path: '/stake', component: Stake },
  { path: '/gov', component: Gov },
  { path: '/faucet', component: Faucet }
]

const router = createRouter({
  history: routerHistory,
  routes
})

export default router
