<template>
    <div>
        <div
            class="w-full flex justify-between items-center gap-4 dark:text-white text-2xl lg:p-4 lg:bg-devbook-light-shades lg:dark:bg-devbook-dark-shades">

            <div @click="navigateToPage('/')" class="w-2/4 flex items-center justify-start">
                <div class="text-xl font-bold cursor-pointer">DEVBOOK</div>
            </div>

            <div class="w-2/4 flex items-center justify-end">
                <div v-if="!authenticated" class="flex flex-col">
                    <button @click="navigateToPage('/login')"
                        class="px-4 py-1 rounded-md text-xl bg-black bg-opacity-0 hover:bg-opacity-10">
                        Login
                    </button>
                    <button @click="navigateToPage('/register')" class="rounded-md text-xs">
                        Register
                    </button>
                </div>
                <button v-else @click="authenticated ? logout() : null" class="rounded-md text-xs">
                    Logout
                </button>

            </div>



        </div>
    </div>
</template>
<script lang="ts">
import { Component, Vue, toNative } from 'vue-facing-decorator';
import store from '../app/store/store';
import router from '../app/router/router';
@Component({})
class TopBar extends Vue {

    navigateToPage(page: string) {
        this.$router.push(page);
    }

    get authenticated() {
        return store.state.user.authenticated;
    }

    logout() {
        store.dispatch('user/LogoutAction')
        router.push('/')
    }

}
export default toNative(TopBar)
</script>