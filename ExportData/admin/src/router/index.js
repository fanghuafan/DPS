import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/register',
    component: () => import('@/views/register/index'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/403',
    component: () => import('@/views/403'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: 'Dashboard', icon: 'dashboard' }
    }]
  },

  {
    path: '/task',
    component: Layout,
    redirect: '/task/list',
    name: 'Task',
    meta: { title: 'Task', icon: 'example' },
    children: [
      {
        path: 'taskList',
        name: 'Task List',
        component: () => import('@/views/task/list/index'),
        meta: { title: 'Task List', icon: 'table' }
      },
      {
        path: 'createTask',
        name: 'Create Task',
        component: () => import('@/views/task/create/index'),
        meta: { title: 'Create Task', icon: 'tree' }
      }
    ]
  },

  {
    path: '/process',
    component: Layout,
    redirect: '/nested/menu1',
    name: 'Process',
    meta: {
      title: 'Process',
      icon: 'nested'
    },
    children: [
      {
        path: 'processList',
        component: () => import('@/views/process/list/index'), // Parent router-view
        name: 'Process List',
        meta: { title: 'Process List', icon: 'nested' },
      },
      {
        path: 'createProcess',
        component: () => import('@/views/process/create/index'),
        meta: { title: 'Process Create', icon: 'form' }
      }
    ]
  },


  {
    path: '/dataBase',
    component: Layout,
    name: 'DataBase',
    meta: {
      title: 'DataBase',
      icon: 'nested'
    },
    children: [
      {
        path: 'dataBaseList',
        name: 'DataBase List',
        component: () => import('@/views/db/list/index'),
        meta: { title: 'Data List', icon: 'form' }
      },
      {
        path: 'dataBaseCreate',
        name: 'DataBase Create',
        component: () => import('@/views/db/create/index'),
        meta: { title: 'DataBase Create', icon: 'form' }
      }
    ]
  },


  {
    path: '/manage',
    component: Layout,
    redirect: '/nested/menu1',
    name: 'Manage System',
    meta: {
      title: 'Manage System',
      icon: 'nested'
    },
    children: [
      {
        path: 'commandPanel',
        name: 'Command Panel',
        component: () => import('@/views/manage/command/index'),
        meta: { title: 'Command Panel', icon: 'form' }
      },
      {
        path: 'password',
        name: 'password',
        component: () => import('@/views/manage/password/index'),
        meta: { title: 'password', icon: 'form' }
      },
      {
        path: 'list',
        name: 'User List',
        component: () => import('@/views/manage/userlist/index'),
        meta: { title: 'User List', icon: 'form' }
      },
      {
        path: 'auth',
        name: 'Auth Center',
        component: () => import('@/views/manage/auth/index'),
        meta: { title: 'Auth Center', icon: 'form' }
      },
      {
        path: 'log',
        name: 'System Log',
        component: () => import('@/views/manage/log/index'),
        meta: { title: 'System Log', icon: 'form' }
      },
    ]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
