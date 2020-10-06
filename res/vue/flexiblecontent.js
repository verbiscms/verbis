if (this.fields === undefined) {
	let temp = {};
	for (const layout in this.getLayouts) {
		if (this.getLayouts[layout] !== undefined) {
			temp[layout] = {};
			const fields = this.getLayouts[layout]['sub_fields'];
			if (fields !== undefined) {
				for (const field in fields) {
					temp[layout][field] = ""
				}
			} else {
				console.log("in fields undef")
			}
		} else {
			console.log("in")
		}
	}
	console.log(temp)
	return temp
} else {
	return this.fields
}


<div class="repeater-item" v-for="(group, groupIndex) in getLayouts" :key="groupIndex">


<div v-if="getLayouts[groupIndex]">
	<div  class="test" v-for="(layout, layoutIndex) in group['sub_fields']" :key="layoutIndex">

	<FieldText v-if="layout.type === 'text'" :layout="layout" :fields.sync="fields[groupIndex][layoutIndex]"></FieldText>

</div>
</div>
</div>