<template>
	<div>
		<div class="container">
			<div v-if="address">
			<SpTokenSend :address="address" />
			<SpTransferList :address="address" />
			</div>
			<div v-else>
				Welcome
			</div>
		</div>
	</div>
</template>

<style scoped>
.container {
	max-width: 800px;
	margin: 2rem auto;
	padding: 1rem;
}
</style>
<script>
export default {
	name: 'Index',
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		}
	}
}
</script>
