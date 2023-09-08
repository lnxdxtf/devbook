<template>
  <div>
    <div class="w-screen h-screen flex flex-col bg-gray-200 dark:bg-zinc-900 p-2">
      <TopBar class="w-full hidden lg:block" />
      <div class="w-full h-full flex">
        <Sidebar />
        <div class="h-full w-full flex justify-center p-4 rounded-md">
          <Main class="w-3/5"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Sidebar from './components/Sidebar.vue'
import TopBar from './components/TopBar.vue'
import Main from './components/Main.vue'
export default {
  components: { Sidebar, TopBar, Main }

}
</script>
