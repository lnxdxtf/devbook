import { App } from "vue";
import router from "./app/router/router";
import store from "./app/store/store";
import { LoginFormDevBookAPI } from "./app/modules/user/interfaces";

async function settings(app: App<Element>) {
    app.use(store)
    app.use(router)

    await store.dispatch('user/LoginAction', { email: "xxx", password: "12283" } as LoginFormDevBookAPI)
}

export default settings