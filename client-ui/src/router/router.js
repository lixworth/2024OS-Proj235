import { createRouter, createWebHashHistory } from 'vue-router';
import { loadingBarApiRef } from '../store/loading';
import { firstLoadingRef } from '../store/first_loading';
import LoginPage from "../view/LoginPage.vue";

const routes = [
    {
        path: '/',
        name: 'index',
        component: LoginPage,
        meta: {
            auth: false
        }
    },
    // NotFoundRoute
];

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
});
router.beforeEach(async (to, from, next) => {
    // const update_user = async () => {
    //     const user_info = useUserInfoStore();
    //     let info = await user_info.update_info();
    //     firstLoadingRef.value = false;
    // };
    if (!from || to.name !== from.name) {
        if (loadingBarApiRef.value) {
            loadingBarApiRef.value.start();
        }
    }
    // let meta = to.meta;
    // if (meta.auth) {
    //     if (!token.hasToken()) {
    //         // @ts-ignore
    //         next({ name: 'login', query: { redirect_url: to.name } });
    //         return;
    //     } else {
    //         await update_user();
    //     }
    // } else {
    //     if (token.hasToken()) {
    //         await update_user();
    //     }
    // }
    firstLoadingRef.value = false;
    next();
});

router.afterEach(() => {});

export default router;