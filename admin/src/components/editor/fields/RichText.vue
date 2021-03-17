<!-- =====================
	Field - Rich Text
	===================== -->
<template>
	<div class=richtext v-if="editor">
		<!-- =====================
			Menu Bubble
			===================== -->
		<editor-menu-bubble :editor="editor" @hide="hideLinkMenu" v-slot="{ commands, isActive, getMarkAttrs, menu }">
			<div class="richtext-bubble" :class="{ 'is-active': menu.isActive }" :style="`left: ${menu.left}px; bottom: ${menu.bottom}px;`">
				<!-- Bold -->
				<button v-if="getByElement('bold')" class="richtext-button richtext-button-first" :class="{ 'is-active': isActive.bold() }" @click="commands.bold">
					<i class="fal fa-bold"></i>
				</button>
				<!-- Italic -->
				<button v-if="getByElement('italic')" class="richtext-button" :class="{ 'is-active': isActive.italic() }" @click="commands.italic">
					<i class="fal fa-italic"></i>
				</button>
				<!-- Strikethrough -->
				<button v-if="getByElement('strike')" class="richtext-button" :class="{ 'is-active': isActive.strike() }" @click="commands.strike">
					<i class="fal fa-strikethrough"></i>
				</button>
				<!-- Underline -->
				<button v-if="getByElement('underline')" class="richtext-button" :class="{ 'is-active': isActive.underline() }" @click="commands.underline">
					<i class="fal fa-underline"></i>
				</button>
				<!-- Link -->
				<div v-if="getByElement('link')">
					<form class="richtext-form" v-if="linkMenuIsActive" @submit.prevent="setLinkUrl(commands.link, linkUrl)">
						<input class="richtext-input" type="text" v-model="linkUrl" placeholder="https://" ref="linkInput" @keydown.esc="hideLinkMenu"/>
						<button class="richtext-button" @click="setLinkUrl(commands.link, null)" type="button">
							<i class="fal fa-times"></i>
						</button>
					</form>
					<template v-else>
						<button class="richtext-button" @click="showLinkMenu(getMarkAttrs('link'))" :class="{ 'is-active': isActive.link() }">
							<i class="fal fa-link"></i>
						</button>
					</template>
				</div>
				<!-- Colour -->
				<div v-if="getByElement('color')">
					<button class="richtext-button" :class="{ 'is-active': isActive.textcolor({ color: 'red' }) }" @click="showTextColorPicker = !showTextColorPicker">
						<i v-if="!showTextColorPicker" class="fal fa-palette"></i>
						<i v-else class="fal fa-times"></i>
					</button>
				</div>
				<!-- Code -->
				<button v-if="getByElement('code')" class="richtext-button richtext-button-last" :class="{ 'is-active': isActive.code() }" @click="commands.code">
					<i class="fal fa-code"></i>
				</button>
			</div>
		</editor-menu-bubble>
		<!-- =====================
			Menu Bar
			===================== -->
		<editor-menu-bar :editor="editor" v-slot="{ commands, isActive, focused }">
			<div>
				<!-- =====================
					Insert Media Modal
					===================== -->
				<Modal :show.sync="showImageModal" class="modal-full-width modal-hide-close" :editor="editor">
					<template slot="text">
						<Uploader :rows="3" :modal="true" :filters="false" class="media-modal" @insert="insertMedia($event, commands.image)" :options="true">
							<template slot="close">
								<button class="btn btn-margin-right btn-icon-mob" @click.prevent="showImageModal = false">
									<i class="feather feather-x"></i>
									<span>Close</span>
								</button>
							</template>
						</Uploader>
					</template>
				</Modal>
				<div class="richtext-menu" :class="{ 'is-focused fadeInDownXs': focused }">
					<!-- Paragraph -->
					<button v-if="getByElement('paragraph')" class="richtext-button" :class="{ 'is-active': isActive.paragraph() }" @click="commands.paragraph">
						<i class="fal fa-paragraph"></i>
					</button>
					<!-- H1 -->
					<button v-if="getByElement('h1')" class="richtext-button" :class="{ 'is-active': isActive.heading({ level: 1 }) }" @click="commands.heading({ level: 1 })">
						<span>H1</span>
					</button>
					<!-- H2 -->
					<button v-if="getByElement('h2')" class="richtext-button" :class="{ 'is-active': isActive.heading({ level: 2 }) }" @click="commands.heading({ level: 2 })">
						<span>H2</span>
					</button>
					<!-- H3 -->
					<button v-if="getByElement('h3')" class="richtext-button" :class="{ 'is-active': isActive.heading({ level: 3 }) }" @click="commands.heading({ level: 3 })">
						<span>H3</span>
					</button>
					<button v-if="getByElement('h4')" class="richtext-button" :class="{ 'is-active': isActive.heading({ level: 4 }) }" @click="commands.heading({ level: 4 })">
						<span>H4</span>
					</button>
					<!-- H5 -->
					<button v-if="getByElement('h5')" class="richtext-button" :class="{ 'is-active': isActive.heading({ level: 5 }) }" @click="commands.heading({ level: 5 })">
						<span>H5</span>
					</button>
					<!-- H6 -->
					<button v-if="getByElement('h6')" class="richtext-button" :class="{ 'is-active': isActive.heading({ level: 6 }) }" @click="commands.heading({ level: 6 })">
						<span>H6</span>
					</button>
					<!-- Unordered List -->
					<button v-if="getByElement('ul')" class="richtext-button" :class="{ 'is-active': isActive.bullet_list() }" @click="commands.bullet_list">
						<i class="fal fa-list-ul"></i>
					</button>
					<!-- Ordered List -->
					<button v-if="getByElement('ol')" class="richtext-button" :class="{ 'is-active': isActive.ordered_list() }" @click="commands.ordered_list">
						<i class="fal fa-list-ol"></i>
					</button>
					<!-- Blockquote -->
					<button v-if="getByElement('blockquote')" class="richtext-button" :class="{ 'is-active': isActive.blockquote() }" @click="commands.blockquote">
						<i class="fal fa-quote-right"></i>
					</button>
					<!-- HR -->
					<button v-if="getByElement('hr')" class="richtext-button" :class="{ 'is-active': isActive.horizontal_rule() }" @click="commands.horizontal_rule">
						<span class="richtext-button-hr"></span>
					</button>
					<!-- Table -->
					<div v-if="getByElement('table')">
						<button class="richtext-button"
								@click="commands.createTable({rowsCount: 3, colsCount: 3, withHeaderRow: false })">
							<i class="fal fa-table"></i>
						</button>
					</div>
					<!-- Undo -->
					<button v-if="getByElement('history')" class="richtext-button" @click="commands.undo">
						<i class="fal fa-undo"></i>
					</button>
					<!-- Redo -->
					<button v-if="getByElement('history')" class="richtext-button" @click="commands.redo">
						<i class="fal fa-redo"></i>
					</button>
					<!-- Colour Picker -->
					<div v-if="showTextColorPicker" class="richtext-colorpicker">
						<color-picker v-model="textColor" @cancel="showTextColorPicker = false" :palette="palette" @input="applyTextColor(commands)"/>
					</div>
					<!-- Code Block -->
<!--					<button v-if="getByElement('code_block')" class="richtext-button" :class="{ 'is-active': isActive.code_block() }" @click="commands.code_block">-->
<!--						<i class="fal fa-code"></i>-->
<!--					</button>-->
					<button v-if="getByElement('code_view')" class="richtext-button" @click="changeCodeView">
						<i class="fal fa-code"></i>
					</button>
					<button class="richtext-button" @click="showImageModal = true">
						<i class="fal fa-image"></i>
					</button>
				</div><!-- /Menu -->
				<!-- =====================
					Table
					===================== -->
				<div v-if="isActive.table()" class="richtext-table">
					<button @click="commands.deleteTable">Delete Table</button>
					<button @click="commands.addColumnBefore">+Col Before</button>
					<button @click="commands.addColumnAfter">+Col After</button>
					<button @click="commands.deleteColumn">Delete Col</button>
					<button @click="commands.addRowBefore">+Row Before</button>
					<button @click="commands.addRowAfter">+Row After</button>
					<button @click="commands.deleteRow">Delete Row</button>
					<button @click="commands.toggleCellMerge">Combine</button>
				</div>
			</div>
		</editor-menu-bar>
		<!-- =====================
			Content
			===================== -->
		<div class="richtext-content" :style="{ 'height' : getEditorHeight }">
			<prism-editor class="richtext-prism prism" v-show="codeView" v-model="code" :highlight="highlighter" line-numbers></prism-editor>
			<editor-content :editor="editor" v-show="!codeView" />
		</div>
		<!-- Message -->
		<transition name="trans-fade-height">
			<span class="field-message field-message-warning" v-if="errors.length">{{ errors[0] }}</span>
		</transition><!-- /Message -->
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->

<script>

import {highlight, languages} from 'prismjs/components/prism-core';
import 'prismjs/components/prism-markup';
import {fieldMixin} from "@/util/fields/fields"

import {Editor, EditorContent, EditorMenuBar, EditorMenuBubble} from 'tiptap'
import {
	VerbisBlockquote,
	VerbisBold,
	VerbisBulletList,
	VerbisCode,
	VerbisCodeBlock,
	VerbisCodeBlockHighlight,
	VerbisColour,
	VerbisHardBreak,
	VerbisHeading,
	VerbisHorizontalRule,
	VerbisItalic,
	VerbisLink,
	VerbisListItem,
	VerbisOrderedList,
	VerbisStrike,
	VerbisTable,
	VerbisUnderline
} from '../../../extensions/tiptap/index'
import {Highlight, History, Image, Search, TableCell, TableHeader, TableRow, TrailingNode,} from 'tiptap-extensions';
import Modal from "@/components/modals/General";
import Uploader from "@/components/media/Uploader";

const Chrome = require('vue-color/src/components/Compact.vue').default;

export default {
	name: "FieldRichText",
	mixins: [fieldMixin],
	props: {
		updating: {
			type: Boolean,
			default: false,
		},
	},
	components: {
		Uploader,
		Modal,
		EditorContent,
		EditorMenuBar,
		EditorMenuBubble,
		'color-picker': Chrome,
	},
	data: () => ({
		html: '',
		editor: false,
		config: {},
		linkMenuIsActive: false,
		showTextColorPicker: false,
		palette: false,
		textColor: '#000000',
		code: '',
		codeView: false,
		showImageModal: false,
		charPosition: 0,
	}),
	mounted() {
		this.config = this.getEditorConfig;
		this.setUpEditor();
		this.setColourPalette();
	},
	watch: {
		code: function (val) {
			this.field = val;
		},
		updating: function () {
			this.$nextTick(() => {
				this.editor.setContent(this.getValue);
			});
		}
	},
	computed: {
		getEditorConfig() {
			return this.$store.state.theme.editor
		},
		getEditorHeight() {
			if ('height' in this.getOptions) {
				return this.getOptions['height'];
			}
			return "400px"
		},
		/*
		 * field()
		 * Replaces <p> tags and strips.
		 * Fire's back up to the parent.
		 */
		field: {
			get() {
				return this.getValue === '<p></p>' ? '' : this.getValue;
			},
			set(value) {
				this.$emit("update:fields", this.getFieldObject(value));
			}
		}
	},
	methods: {
		/*
		 * validate()
		 * Fires when the publish button is clicked.
		 */
		validate() {
			this.errors = [];
			this.validateRequired();
		},
		insertMedia(media, command) {
			const src = media.url;
			if (src !== null) {
				command({ src })
			}
			this.showImageModal = false;
		},
		setContent(content) {
			if (this.editor) this.editor.setContent(content);
		},
		setUpEditor() {
			const extensions = this.createExtensions();
			this.setDefaultValue();
			this.$nextTick(() => {
				this.editor = new Editor({
					content: this.field,
					onUpdate: ({ getHTML }) => {
						this.errors = [];
						this.html = getHTML();
						if (this.html === '<p></p>') this.html = '';
						this.field = this.html;
					},
					onTransaction: ({ state }) => {
						this.charPosition = state.selection.anchor;
					},
					extensions: extensions,
				});
			})
		},
		highlighter(code) {
			return highlight(code, languages.html);
		},
		getByElement(element) {
			if (!this.config) {
				return true;
			} else {
				const modules = this.config.modules;
				if (modules.includes(element)) {
					return modules[modules.indexOf(element)];
				}
			}
			return false;
		},
		processNode(module) {
			return module === undefined ? false : module;
		},
		hideLinkMenu() {
			this.linkUrl = null;
			this.linkMenuIsActive = false;
		},
		setLinkUrl(command, url) {
			command({href: url});
			this.field = this.html;
			this.hideLinkMenu()
		},
		showLinkMenu(attrs) {
			this.linkUrl = attrs.href;
			this.linkMenuIsActive = true;
			this.$nextTick(() => {
				this.$refs.linkInput.focus()
			})
		},
		setColourPalette() {
			this.palette = this.config.options.palette.length > 0 ? this.config.options.palette : false
		},
		applyTextColor(commands) {
			const { r, g, b, a } = this.textColor.rgba;
			commands.textcolor({ color: `rgba(${r},${g},${b},${a})` })
			this.showTextColorPicker = false
		},
		changeCodeView() {
			if (this.codeView) {
				this.editor.setContent(this.code)
				this.field = this.code;
			} else {
				this.code = this.field;
			}
			this.codeView = !this.codeView;
		},
		// TODO Extra spaces are coming back
		process(str) {
			let div = document.createElement('div');
			div.innerHTML = str.trim();
			return this.format(div, 0).innerHTML;
		},
		format(node, level) {
			let indentBefore = new Array(level++ + 1).join('  '),
				indentAfter  = new Array(level - 1).join('  '),
				textNode;

			for (let i = 0; i < node.children.length; i++) {

				textNode = document.createTextNode('\n' + indentBefore);
				node.insertBefore(textNode, node.children[i]);

				this.format(node.children[i], level);

				if (node.lastElementChild === node.children[i]) {
					textNode = document.createTextNode('\n' + indentAfter);
					node.appendChild(textNode);
				}
			}
			return node;
		},
		createExtensions() {
			let extensions = [],
				headingArr = [];

			extensions.push(new Image())

			this.config.modules.forEach(module => {
				let node = this.processNode(this.config.options[module]);

				switch (module) {
					// Nodes
					case 'blockquote':
						extensions.push(new VerbisBlockquote(node));
						break;
					case 'ul':
						extensions.push(new VerbisBulletList(node));
						break;
					case 'code_block':
						extensions.push(new VerbisCodeBlock(node));
						break;
					case 'code_block_highlight':
						extensions.push(new VerbisCodeBlockHighlight(node));
						break;
					case 'hardbreak':
						extensions.push(new VerbisHardBreak(node));
						break;
					case 'h1':
						headingArr.push(1);
						break;
					case 'h2':
						headingArr.push(2);
						break;
					case 'h3':
						headingArr.push(3);
						break;
					case 'h4':
						headingArr.push(4);
						break;
					case 'h5':
						headingArr.push(5);
						break;
					case 'h6':
						headingArr.push(6);
						break;
					case 'hr':
						extensions.push(new VerbisHorizontalRule(node));
						break;
					case 'ol':
						extensions.push(new VerbisOrderedList(node));
						break;
					// Table
					case 'table':
						extensions.push(new VerbisTable(node));
						extensions.push(new TableHeader());
						extensions.push(new TableCell());
						extensions.push(new TableRow());
						break;
					// Marks
					case 'bold':
						extensions.push(new VerbisBold(node));
						break;
					case 'code':
						extensions.push(new VerbisCode(node));
						break;
					case 'italic':
						extensions.push(new VerbisItalic(node));
						break;
					case 'link':
						extensions.push(new VerbisLink(node));
						break;
					case 'strike':
						extensions.push(new VerbisStrike(node));
						break;
					case 'underline':
						extensions.push(new VerbisUnderline(node));
						break;
					// Extensions
					case 'color':
						extensions.push(new VerbisColour());
						break;
					// case 'collaboration':
					//     extensions.push(new Collaboration());
					//     break;
					// case 'focus':
					//     extensions.push(new Focus());
					//     break;
					case 'history':
						extensions.push(new History());
						break;
					// case 'placeholder':
					//     extensions.push(new Placeholder());
					//     break;
					case 'search':
						extensions.push(new Search());
						break;
					case 'trailing_node':
						extensions.push(new TrailingNode());
						break;
					// Plugins
					case 'highlight':
						extensions.push(new Highlight());
						break;
				}
			});

			// Add heading if any headings exist
			if (headingArr.length > 0) {
				extensions.push(new VerbisHeading(headingArr, {
					'1': this.processNode(this.config.options['h1']),
					'2': this.processNode(this.config.options['h2']),
					'3': this.processNode(this.config.options['h3']),
					'4': this.processNode(this.config.options['h4']),
					'5': this.processNode(this.config.options['h5']),
					'6': this.processNode(this.config.options['h6']),
				}));
			}

			// Add list item if ol or ul exists
			if (this.config.modules.includes('ul') || this.config.modules.includes('ol')) {
				let node = this.processNode(this.config.options['li']);
				extensions.push(new VerbisListItem(node));
			}

			return extensions;
		},
	},
	beforeDestroy() {
		this.editor.destroy()
	},
}
</script>

<!-- =====================
	Styles
	===================== -->
<style lang="scss">

// Variables
$richtext-border-radius: 10px;

.richtext {
	$self: &;

	// Buttons
	// =========================================================================

	&-menu {
		position: relative;
		display: flex;
		flex-wrap: nowrap;
		overflow-x: scroll;
		overflow: visible;
		width: 100%;
		background-color: $white;
		border-top-left-radius: $richtext-border-radius;
		border-top-right-radius: $richtext-border-radius;
		border: 1px solid $grey-light;
		border-bottom-width: 0;
	}

	// Buttons
	// =========================================================================

	&-button {
		display: flex;
		justify-content: center;
		align-items: center;
		border: none;
		background-color: $white;
		width: 50px;
		min-width: 40px;
		height: 42px;
		border-right: 1px solid $grey-light;
		color: $grey;
		cursor: pointer;
		outline: none;

		i,
		span {
			font-size: 16px;
			font-weight: 500;
			transition: color 200ms ease;
			will-change: color;
		}

		span {
			font-size: 15px;
		}

		&:first-child {
			border-top-left-radius: $richtext-border-radius;
		}

		&:hover {

			i,
			span {
				color: $primary;
			}
		}

		&-hr {
			display: block;
			width: 60%;
			height: 2px;
			background-color: $grey;
			border-radius: 10px;
		}
	}

	// Bubble
	// =========================================================================

	&-bubble {
		position: absolute;
		display: flex;
		z-index: 9999;
		background: $white;
		border-radius: 5px;
		margin-bottom: .5rem;
		transform: translateX(-50%);
		border: 1px solid $grey-light;
		overflow: visible;
		visibility: hidden;
		opacity: 0;
		transition: opacity .2s, visibility .2s;

		&.is-active {
			opacity: 1;
			visibility: visible;
		}

		#{$self}-button {
			border-radius: 0;
			width: 36px;
			height: 36px;

			i,
			span {
				font-size: 14px;
			}

			&-last {
				border-right: none;
			}
		}
	}

	// Form
	// =========================================================================

	&-form {
		display: flex;

		input {
			border: none;
			font-size: 14px;
			color: $copy-light;
			padding: 4px 10px;
			outline: none;
		}
	}

	// Colour Picker
	// =========================================================================

	&-colorpicker {
		position: absolute;
		display: flex;
		flex-direction: column;
		align-items: flex-end;
		top: 20px;
		right: 20px;
		background-color: $white;
		border: 1px solid $grey-light;
		z-index: 99;
		padding: 10px;
		border-radius: 4px;

		.vc-compact {
			box-shadow: none;
			width: auto;
		}
	}

	// Table
	// =========================================================================

	&-table {
		position: relative;
		display: flex;
		align-items: center;
		justify-content: space-between;
		width: 100%;
		background-color: $white;
		border: 1px solid $grey-light;
		border-bottom-width: 0;
		padding: 0 2px;

		button {
			background: transparent;
			border: none;
			font-size: 12px;
			color: $grey;
			padding: 8px 0;
			text-align: center;
			border-right: 1px solid $grey-light;
			flex-grow: 2;
			outline: none;
			cursor: pointer;
			transition: color 200ms ease;
			will-change: color;

			&:last-child {
				border: none;
			}

			&:hover {
				color: $primary;
			}
		}
	}

	// Code
	// =========================================================================

	&-prism {
		position: relative;
		top: 0;
		left: 0;
		width: 100% !important;
		height: 100% !important;
		border: none;
		font-size: 14px;
		resize: none;
	}

	// Content (Body)
	// =========================================================================

	&-content {
		position: relative;
		background-color: $white;
		border: 1px solid $grey-light;
		border-bottom-left-radius: $richtext-border-radius;
		border-bottom-right-radius: $richtext-border-radius;
		height: 400px;
		min-height: 400px;
		padding: 20px;
		overflow-y: scroll;

		& > div,
		.ProseMirror {
			min-height: 400px;
		}

		* {
			outline: none;
		}

		// Body Styles
		// =========================================================================

		h1,
		h2,
		h3,
		h4,
		h5,
		h6 {
			color: $secondary;
		}

		strong {
			font-weight: 900;
		}

		h1 {
			font-size: 46px;
			margin-bottom: 10px;
		}

		h2 {
			font-size: 34px;
			margin-bottom: 10px;
		}

		h3 {
			font-size: 28px;
			margin-bottom: 10px;
		}

		h4 {
			font-size: 22px;
			margin-bottom: 10px;
		}

		h5 {
			font-size: 18px;
		}

		h6 {
			font-size: 14px;
			text-transform: initial;
		}

		a {
			color: $primary;
			text-decoration: underline;
			font-weight: bold;
		}

		ul,
		ol {
			margin: 0 0 0 10px;
			padding: 0;
			list-style-position: outside;

			li {
				line-height: 1;
			}

			p {
				display: inline-block;
				margin-bottom: 0;
			}
		}

		ul {
			list-style: none;

			li {
				display: block;

				&:before {
					position: relative;
					content: "\2022";
					color: $secondary;
					font-weight: bold;
					display: inline-block;
					width: 1em;
				}
			}
		}

		ol {
			list-style-type: decimal;
			list-style-position: inside;
		}

		table {

			tr,
			th,
			td {
				border: 1px solid $grey-light;
			}


			td {
				min-width: 50px;
				padding: 5px;
			}
		}

		blockquote {
			display: flex;
			padding: 0;
			margin: 0;
			font-style: italic;

			&:before,
			&:after {
				display: inline-block;
				content: "\"";
				color: $grey;
			}

			&:before {
				margin-right: 2px;
			}

			&:after {
				margin-left: 2px;
			}
		}

		.code {
			background: red;
		}
	}
}
</style>