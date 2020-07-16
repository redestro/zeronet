/* eslint-disable */
<template>
	<td class="cell" @click="strike">{{ mark }}</td>
</template>

<script>
import '@/assets/_grid.scss';
	export default {
		props: ['name'],
		data () {
			return {
 				frozen: false,
				mark: ''
			}	
		},
		methods: {
			strike () {
				if (! this.frozen) {
					this.mark = this.$parent.activePlayer
					this.frozen = true
					Event.$emit('strike', this.name)
				}
			}
		},
		created () {
			Event.$on('clearCell', () => {
				this.mark = ''
				this.frozen = false
			})
			Event.$on('freeze', () => this.frozen = true)
		}
	}
</script>
