<template>
    <div>
        <div class="w-full h-full flex justify-center items-center text-black dark:text-white input-devbook-login">
            <div class="w-2/5 flex p-2">

                <div class="w-full flex flex-col gap-2 p-4 bg-white dark:bg-zinc-700 rounded-md">
                    <div class="w-full p-4 flex items-center justify-center">
                        <span class="text-4xl font-bold">Login</span>
                    </div>
                    <div class="w-full flex flex-col">
                        <span>Email</span>
                        <input class="text-black" type="text">
                    </div>
                    <div class="w-full flex flex-col">
                        <span>Password</span>
                        <input class="text-black" type="text">
                    </div>
                    <div class="w-full flex justify-center p-4">
                        <div class="w-3/4 p-2 flex justify-center items-center bg-devbook-light-accent dark:bg-devbook-dark-accent rounded-md cursor-pointer hover:hue-rotate-15 ease-linear duration-150">
                            <span class="text-2xl text-white">Login</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
export default {

}
</script>