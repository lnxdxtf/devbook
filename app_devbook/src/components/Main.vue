<template>
    <div>
        <div class="w-full h-full lg:h-[86vh] overflow-y-auto p-4 rounded-md bg-devbook-light-shades dark:bg-devbook-dark-shades">
            <router-view v-slot="{ Component }">
                <component :is="Component" class="w-full h-full animate__animated animate__fadeIn" />
            </router-view>
        </div>
    </div>
</template>
<script>
export default {

}
</script>