<template>
    <div>
        <div
            class="w-3/4 h-full flex flex-col justify-around items-center p-2 rounded-md text-black dark:text-white bg-devbook-light-shades dark:bg-devbook-dark-shades">

            <template v-if="randomUsers">
                <div v-for="user in randomUsers" :key="user.id">
                    <div class="w-full flex flex-col items-center text-black dark:text-white p-2 gap-2">
                        <img :src="`${devbookCDNUserImg}${user.id}.png`" alt="user image" class="rounded-full w-12">
                        <span class="font-bold text-sm">@{{ user.nick }}</span>
                        <div @click="followUser(user.id)"
                            class="text-white bg-devbook-main-one hover:bg-devbook-main-two cursor-pointer font-bold p-1 rounded-md ease-in-out duration-300">
                            FOLLOW
                        </div>
                    </div>
                </div>
            </template>

        </div>
    </div>
</template>
<script lang="ts">
import { Component, Vue, toNative } from 'vue-facing-decorator';
import router from '../app/router/router';
import store from '../app/store/store';
import { NotificationError } from '../app/utils/notification';
import cdn_paths from '../cdn_paths.json'

@Component({})
class UsersSuggetions extends Vue {
    get randomUsers() {
        return store.state.user.randomUsers
    }
    get authenticated() {
        return store.state.user.authenticated
    }

    get devbookCDNUserImg() {
        return `${store.state.devbookCDN}${cdn_paths.imgs.user.profile}/user_`
    }

    async followUser(id: number) {
        if (!this.authenticated) {
            NotificationError('You need to be logged in to follow users')
            router.push('/login')
            return
        }
        store.dispatch('user/FollowUserAction', id)
        return

    }
}
export default toNative(UsersSuggetions)
</script>
