import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/Index.vue'
import Types from '@/views/Types.vue'
import Address from '@/views/Address.vue'
import Block from '@/views/Block.vue'
import Blocks from '@/views/Blocks.vue'
import Send from '@/views/Send.vue'
import Wallet from '@/views/Wallet.vue'
import Legacy from '@/views/Legacy.vue'

const routerHistory = createWebHistory()
const routes = [
	{
		path: '/',
		component: Index
	},
	{ path: '/types', component: Types },
]

const router = createRouter({
	history: routerHistory,
	routes
})

export default router
