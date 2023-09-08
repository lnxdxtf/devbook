<template>
   <div>

      <SideBarMobile class="lg:hidden block h-full" />
      <SideBarDesktop class="lg:block hidden h-full" />

   </div>
</template>
<script>
import SideBarMobile from './mobile/SideBarMobile.vue'
import SideBarDesktop from './desktop/SideBarDesktop.vue'
import { mapState } from 'vuex';
export default {
   components: { SideBarMobile, SideBarDesktop },
   computed: {
      ...mapState({
         isMobile: state => state.isMobile
      })
   }
}
</script>
