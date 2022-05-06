<template>
<div>
  <Suspense><SpSystemBar /></Suspense>
  <div class="navbar-wrapper">
    <div class="navbar-section">
      <slot name="logo">
        <router-link
          :to="'/'"
          class="sp-nav-link hide-on-small"
          :alt="'Home'"
          :title="'Home'"
        >
          <div style="display: flex; align-items: center">
            <img
              width="40"
              height="40"
              src="/public/defund.svg"
            >
            <span
              style="
                padding: 4px 8px;
                margin-left: 10px;
                background: rgba(0, 0, 0, 0.03);
                border-radius: 24px;
                font-weight: 500;
                font-size: 10px;
              "
            >
              Alpha
            </span>
          </div>
        </router-link>
      </slot>
      <router-link
        v-for="(link, lid) in links"
        :key="`link-${lid}`"
        :to="link.url"
        class="sp-nav-link"
        :alt="link.name"
        :title="link.name"
      >
        <div :class="link.url === activeRoute ? 'link-active' : ''">
          {{ link.name }}
        </div>
      </router-link>
    </div>
    <div class="navbar-section">
      <SpAcc />
    </div>
  </div>
</div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'

import { SpAcc, SpButton, SpModal, SpSystemBar } from '@starport/vue'

export interface NavbarLink {
  name: string
  url: string
}

export default defineComponent({
  name: 'SpNavbar',

  components: {
    SpAcc,
    SpModal,
    SpButton,
    SpSystemBar
  },

  props: {
    links: {
      type: Object as PropType<NavbarLink[]>,
      required: true
    },
    activeRoute: {
      type: String,
      required: false
    }
  }
})
</script>

<style scoped lang="scss">
.navbar-wrapper {
  display: flex;
  justify-content: space-between;
  height: 80px;
  left: 0;
  right: 0;
  top: 0;
  background: #ffffff;
  margin-bottom: 54px;
}

.navbar-section {
  display: flex;
  padding: 20px;
  align-items: center;
}

.sp-nav-link {
  font-size: 16px;
  line-height: 130%;
  color: rgba(0, 0, 0, 0.667);
  font-weight: 400;
  text-decoration: none;
  cursor: pointer;
  margin: 0 1rem;
  transition: font-weight 0.2s ease, color 0.2s ease;
}

.sp-nav-link:hover {
  opacity: 0.8;
}

.sp-nav-link.selected {
  font-weight: 600;
  color: #000000;
}

.description-grey {
  font-size: 13px;
  line-height: 153.8%;
  color: rgba(0, 0, 0, 0.667);
}

.external-link {
  font-weight: 600;
  font-size: 16px;
  cursor: pointer;
}

.external-link:hover {
  opacity: 0.8;
}

.link-active {
  font-weight: 500;
  color: #000;
}

@media (max-width: 600px) {
  .hide-on-small {
    display: none;
  }
}
</style>
