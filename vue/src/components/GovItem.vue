<template>
    <div class="row">
        <div class="css-1van9nl"><div class="chakra-stack css-ebgwia"><div class="css-acwl5z"><p class="chakra-text css-1lutiry">#ID</p><div class="chakra-skeleton css-143twa6"><div class="css-79elbk">1</div></div></div><div class="css-acwl5z"><p class="chakra-text css-1lutiry">Title</p><div class="chakra-skeleton css-143twa6"><div class="css-79elbk"><span class="css-16vxv8a"><a class="" href="/proposal/1">Minting Param Change</a></span></div></div></div><div class="css-acwl5z"><p class="chakra-text css-1lutiry">Status</p><div class="chakra-skeleton css-143twa6"><div class="css-79elbk"><span class="chakra-badge css-1xazz9g">Passed</span></div></div></div><div class="css-acwl5z"><p class="chakra-text css-1lutiry">Voting Start</p><div class="chakra-skeleton css-143twa6"><div class="css-79elbk">21.04.2022, 03:29:07</div></div></div><div class="css-acwl5z"><p class="chakra-text css-1lutiry">Voting End</p><div class="chakra-skeleton css-143twa6"><div class="css-79elbk">23.04.2022, 03:29:07</div></div></div><div class="css-acwl5z"><p class="chakra-text css-1lutiry">Total Deposit</p><div class="chakra-skeleton css-143twa6"><div class="css-79elbk">100 FETF</div></div></div></div></div>
    </div>
</template>
 
<script>
import VGrid, {VGridVueTemplate} from "@revolist/vue3-datagrid";
import { computed } from 'vue';
import { useStore } from 'vuex';
import { SpTheme, SpButton } from '@starport/vue';
export default {
  name: "DetailButton",
  components: {
    VGrid, SpTheme
  }
};
</script>