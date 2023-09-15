<template>
    <div>
        <div
            class="w-full flex justify-between items-center gap-4 dark:text-white text-2xl lg:p-4 lg:bg-devbook-light-shades lg:dark:bg-devbook-dark-shades">

            <div @click="navigateToPage('/')" class="w-2/4 flex items-center justify-start">
                <div class="text-xl font-bold cursor-pointer">DEVBOOK</div>
            </div>

            <div class="w-2/4 flex items-center justify-end">
                <div class="flex flex-col">
                    <button @click="navigateToPage('/login')" class="px-4 py-1 rounded-md text-xl bg-black bg-opacity-0 hover:bg-opacity-10">
                        Login
                    </button>
                    <button @click="navigateToPage('/register')" class="rounded-md text-xs">
                        Register
                    </button>
                </div>
            </div>



        </div>
    </div>
</template>
<script lang="ts">
import { Component, Vue, toNative } from 'vue-facing-decorator';
@Component({})
class TopBar extends Vue {

    navigateToPage(page: string) {
        this.$router.push(page);
    }

}
export default toNative(TopBar)
</script>