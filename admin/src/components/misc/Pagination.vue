<!-- =====================
	Pagination
	===================== -->
<template>
	<div class="pagination-cont"  v-if="pagination">
		<div class="pagination">
			<!-- Beginning -->
			<div class="pagination-item" :class="{ 'pagination-item-disabled' : !getPagination.prev }" @click="update(1)">
				<i class="feather feather-chevrons-left"></i>
			</div>
			<!-- Previous -->
			<div class="pagination-item" :class="{ 'pagination-item-disabled' : !getPagination.prev }" @click="update(getPagination.page - 1)">
				<i class="feather feather-chevron-left"></i>
			</div>
			<!-- Pages -->
			<div class="pagination-item pagination-item-page" :class="{ 'pagination-item-active' : page === getPagination.page }" v-for="page in getPages" :key="page" @click="update(page)">
				{{ page }}
			</div>
			<!-- Next -->
			<div class="pagination-item" :class="{ 'pagination-item-disabled' : !getPagination.next }" @click="update(getPagination.page + 1)">
				<i class="feather feather-chevron-right"></i>
			</div>
			<!-- End -->
			<div class="pagination-item" :class="{ 'pagination-item-disabled' : !getPagination.next }" @click="update(getPagination.pages)">
				<i class="feather feather-chevrons-right"></i>
			</div>
		</div><!-- /Pagination  -->
		<div>
			{{ pagination }}
			{{ getPages }}
<!--			{{ getPages }}-->
<!--			1 - 1 2 3 4 5-->
<!--			2 - 1 2 3 4 5-->
<!--			3 - 1 2 3 4 5-->
<!--			4 - 2 3 4 5 6-->
<!--			5 - 2 3 4 5 6-->
<!--			6 - 2 3 4 5 6-->
		</div>
	</div><!-- /Pagination Cont -->
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

export default {
	name: "Pagination",
	props: {
		pagination: {
			type: Object,
		}
	},
	data: () => ({
		pages: [1, 2, 3, 4, 5, 6],
		activePage: 0,
	}),
	methods: {
		/*
		 * update()
		 * Init the tabs and push to the tabs array.
		 */
		update(page) {
			console.log(page)
			this.$emit('update', `page=${page}`)
		},
	},
	computed: {
		getPages() {
			let forwardPages = 0;
			let backwardsPages = 0;
			const curPage = this.getPagination.page;

			if ((curPage - 3) < 0) {
				backwardsPages = Math.abs(Math.abs(curPage - 3) - curPage)
			} else {
				backwardsPages = curPage - (curPage - 3)
			}

			if ((curPage + 3) > curPage) {
				let totalMax = this.getPagination.pages - (curPage + 3)
				forwardPages = forwardPages + totalMax
			} else {
				forwardPages = (curPage + 3) - curPage
			}

			let arr = [];
			for (let i = 0; i <= forwardPages; i++) {
				if ((curPage + i) !== 0) {
					arr.push(curPage + i)
				}
			}
			for (let i = 1; i <= backwardsPages; i++) {
				if ((curPage - i) !== 0) {
					arr.push(curPage - i)
				}
			}
			arr.sort();

			return arr;
		},
		getPagination() {
			return this.pagination;
		}
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">

// Variables
$pagination-border-radius: 4px;

.pagination {
	$self: &;

	display: inline-flex;
	justify-content: flex-start;
	border: 1px solid $grey-light;
	background-color: $white;
	border-radius: 4px;
	margin-left: auto;
	margin-right: 0;

	// Container
	// =========================================================================

	&-cont {
		width: 100%;
		display: flex;
		justify-content: flex-end;
	}

	// Item
	// =========================================================================

	&-item {
		cursor: pointer;
		padding: 6px 16px;
		user-select: none;
		transition: background-color 160ms ease, box-shadow 160ms ease, color 160ms ease;
		will-change: background-color, box-shadow, color;
		color: $grey;
		border-radius: 3px;
		font-size: 16px;

		i {
			color: $grey;
			transition: color 160ms ease;
			will-change: color;
		}

		&-active {
			background-color: $primary;
			color: $white;
			box-shadow: 0 0 10px 2px rgba($primary, 0.20);
			transform: scale(1.02);

			i {
				color: $white;
			}
		}

		&-disabled {
			background-color: rgba($secondary, 0.1);
			pointer-events: none;
		}

		&:hover:not(&-active) {
			color: $primary;

			i {
				color: $primary;
			}
		}

		&:first-child {
			border-top-left-radius: $pagination-border-radius;
			border-bottom-left-radius: $pagination-border-radius;
		}

		&:last-child {
			border-right: none;
			border-top-right-radius: $pagination-border-radius;
			border-bottom-right-radius: $pagination-border-radius;
		}
	}
}

</style>