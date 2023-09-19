<template>
  <div>
    <div class="w-screen h-screen flex flex-col bg-gray-200 dark:bg-zinc-900 p-2 hidden-scroll ">
      <TopBar class="w-full hidden lg:block" />
      <div class="w-full h-full flex">
        <Sidebar class="lg:relative absolute" />
        <div class="h-[88vh] w-full flex justify-center p-4 rounded-md">
          <div class="w-full flex justify-center">
            <Main class="w-full lg:w-3/5 m-4" />
            <UsersSuggetions class="w-1/5 m-4" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, toNative } from 'vue-facing-decorator';
import Sidebar from './components/Sidebar.vue'
import TopBar from './components/TopBar.vue'
import Main from './components/Main.vue'
import UsersSuggetions from './components/UsersSuggetions.vue'
@Component({
  components: { Sidebar, TopBar, Main, UsersSuggetions }
})
class App extends Vue { }
export default toNative(App)
</script>
