<script src="../marks/Bold.js"></script>
<template>
    <field-editor ref="field" v-model="value" :options="options" :show-editor="showEditor" @toggleShowEditor="toggleEditor" @cancel="cancel" @resetvalue="resetValue" @change="update">
        <template slot="preview">
      <span v-if="color === null || color === ''" class="dvs-italic">Currently No Value</span>
            <div v-else class="dvs-flex dvs-items-center">
                <div class="dvs-w-4 dvs-h-4 dvs-rounded-full dvs-mr-4" :style="{'background-color': color}"></div>
            </div>
        </template>
        <template slot="editor">
            <color-picker v-model="color" @cancel="cancel"/>
        </template>
    </field-editor>
</template>

<script>
    import Field from '../../../mixins/Field';

    // eslint-disable-next-line no-undef
    const Chrome = require(/* webpackChunkName: "vue-color" */ 'vue-color/src/components/Chrome.vue').default;
    // eslint-disable-next-line no-undef
    const tinycolor = require(/* webpackChunkName: "tinycolor" */ 'tinycolor2');

    export default {
        name: 'ColorEditor',
        components: {
            FieldEditor: () => import(/* webpackChunkName: "devise-editors" */ './Field'),
            'color-picker': Chrome,
        },
        mixins: [Field],
        props: {
            value: {
                type: Object,
                required: true,
            },
            options: {
                type: Object,
                required: true,
            },
        },
        data() {
            return {
                showEditor: false,
                originalValue: null,
            };
        },
        computed: {
            getMaxLength() {
                if (typeof this.settings !== 'undefined' && typeof this.settings.maxlength !== 'undefined') {
                    return this.settings.maxlength;
                }
                return '';
            },
            color: {
                get() {
                    return tinycolor(this.value.color).toHex();
                },
                set(color) {
                    let valueObj = Object.assign(this.value, {color: null});
                    if (color !== null) {
                        valueObj = Object.assign(this.value, {
                            color: `rgba(${color.rgba.r},${color.rgba.g},${color.rgba.b},${color.rgba.a})`,
                        });
                    }
                    this.$emit('input', valueObj);
                    this.$emit('change', valueObj);
                },
            },
        },
        mounted() {
            this.originalValue = this.value.color;
        },
        methods: {
            toggleEditor() {
                this.showEditor = !this.showEditor;
            },
            cancel() {
                const rgba = this.convertColor(this.originalValue);
                this.color = {rgba};
                this.enabled = this.originalValue.enabled;
            },
            convertColor(color) {
                return tinycolor(color).toRgb();
            },
            resetValue() {
                this.enabled = false;
                this.color = null;
            },
        },
    };
</script>
