import { App } from "vue";
import router from "./app/router/router";
import store from "./app/store/store";
import Toast, { POSITION } from "vue-toastification";
import "vue-toastification/dist/index.css";
import { timePassed } from "./app/utils/directives/time";

async function settings(app: App<Element>) {
    // Plugins
    app.use(store)
    app.use(router)
    app.use(Toast, {
        position: POSITION.TOP_RIGHT,
        timeout: 4000,
        transition: "Vue-Toastification__slideBlurred",
    })

    app.directive('timePassed', timePassed)
    
    await userSettings()
    setListeners()
}

function setListeners() {
    window.addEventListener("resize", () => { store.commit("SET_ISMOBILE", document.body.offsetWidth < 1024) })
}

/**
 * Check if user is logged and set the user data in store
 * @returns void
*/

async function userSettings() {
    store.dispatch('user/GetUserLoggedAction')
    await store.dispatch('user/GetRandomUsersAction')
}

export default settings