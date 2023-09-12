import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router'
import routes from './routes.json'

function LazyLoaderPage(pageName: string) {
    return import(`../../pages/${pageName}/${pageName}.vue`)
}

const dynamicRoutes: RouteRecordRaw[] = routes.map(r => {
    return {
        path: r.dynamic ? `${r.path}/:id` : r.path,
        name: r.name,
        component: LazyLoaderPage(r.component),
        meta: r.meta,
        beforeEnter: (to: any, from: any, next: any) => {
            if (r.auth) {
                next()
            } else {
                next()
            }
        }
    }
}) as RouteRecordRaw[]

export default createRouter({
    history: createWebHistory(),
    routes: [...dynamicRoutes, { path: '/:pathMatch(.*)*', redirect: "/", }]
})