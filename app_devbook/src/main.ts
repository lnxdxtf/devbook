import { createApp } from 'vue'
import './app.css'
import 'animate.css';
import App from './App.vue'
import settings from './settings';


async function INIT() {
    const app = createApp(App)
    await settings(app)
    app.mount('#app')
}

await INIT()
