<template>
   <div>

      <SideBarMobile class="lg:hidden block h-full" />
      <SideBarDesktop class="lg:block hidden h-full" />

   </div>
</template>
<script lang="ts">
import SideBarMobile from './mobile/SideBarMobile.vue'
import SideBarDesktop from './desktop/SideBarDesktop.vue'
import { Component, Vue, toNative } from 'vue-facing-decorator';
import store from '../app/store/store';

@Component({
   components: { SideBarMobile, SideBarDesktop },
})
class SideBar extends Vue {
   get isMobile() {
      return store.state.isMobile
   }
}
export default toNative(SideBar)
</script>
