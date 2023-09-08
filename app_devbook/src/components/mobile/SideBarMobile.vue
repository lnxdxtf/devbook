<template>
    <div>
        <div class="w-12 h-12 flex items-center justify-center cursor-pointer rounded-md text-2xl text-white bg-devbook-light-accent dark:bg-devbook-dark-accent"
            type="button" data-drawer-target="drawer-navigation" data-drawer-show="drawer-navigation"
            aria-controls="drawer-navigation">
            <i class="fa-solid fa-bars"></i>
        </div>

        <!-- drawer component -->
        <div id="drawer-navigation"
            class="fixed top-0 left-0 z-40 w-64 h-screen p-4 overflow-y-auto transition-transform -translate-x-full bg-devbook-light-shades dark:bg-devbook-dark-shades"
            tabindex="-1" aria-labelledby="drawer-navigation-label">
            <h5 id="drawer-navigation-label" class="text-base font-semibold text-white">Menu
            </h5>
            <button type="button" data-drawer-hide="drawer-navigation" aria-controls="drawer-navigation"
                class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 absolute top-2.5 right-2.5 inline-flex items-center dark:hover:bg-gray-600 dark:hover:text-white">
               
                <span class="sr-only">Close menu</span>
            </button>
            <div class="py-4 overflow-y-auto">
                
            </div>
        </div>
    </div>
</template>
<script>
export default {

}
</script>
<style lang="">
    
</style>