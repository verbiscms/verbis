<!-- =====================
	Nav Item
	===================== -->
<template>
	<div class="item card card-small-box-shadow card-expand card-expand-full-width">
		<Collapse :show="false" class="collapse-border-bottom">
			<template v-slot:header>
				<div class="card-header">
					<p class="card-title">{{ item['text'] }}</p>
					<div class="card-controls">
						<i class="feather feather-chevron-down"></i>
					</div>
				</div><!-- /Card Header -->
			</template>
			<template v-slot:body @dragstart="test();">
				<div class="card-body">
					<!-- Label -->
					<FormGroup label="Label*">
						<input class="form-input form-input-white" type="text" v-model="item['text']">
					</FormGroup>
					<!-- Title -->
					<FormGroup label="Title">
						<input class="form-input form-input-white" type="text" v-model="item['title']">
					</FormGroup>
					<!-- Rel -->
					<FormGroup label="Rel">
						<input class="form-input form-input-white" type="text" v-model="item['rel']">
					</FormGroup>
					<!-- Open Tab -->
					<FormGroup label="Open link in new tab">
						<div class="toggle">
							<input type="checkbox" class="toggle-switch" id="cache-frontend" checked v-model="item['new_tab']" :true-value="true" :false-value="false" />
							<label for="cache-frontend"></label>
						</div>
					</FormGroup>
				</div><!-- /Card Body -->
			</template>
		</Collapse><!-- /Cache assets? -->
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->
<script>

import Collapse from "@/components/misc/Collapse";
import FormGroup from "@/components/forms/FormGroup";

export default {
	name: "NavItem",
	props: {
		item: {
			type: Object,
		},
	},
	components: {
		Collapse,
		FormGroup,
	},
	data: () => ({
		show: true,
		gridSize: 70,
		isActive: true,
	}),
	methods: {
		test() {
			console.log("fired")
		}
	},
	computed: {
	}
}

</script>

<!-- =====================
	Styles
	===================== -->
<style scoped lang="scss">


.item {
	//background: $white;
	border-radius: 4px;
	margin-bottom: 1rem;
	width: 50%;

	.card-header {
		padding: 0.8rem;
	}

	&-cont {
		margin-left: 50px;
	}
	.card {

		.card-header {
			padding: 1rem;
		}
	}
}

</style>
