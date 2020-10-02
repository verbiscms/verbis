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
					<!-- H4 -->
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
		<div class="richtext-content" :class="{ 'richtext-content-codeview' : codeView }" >
			<textarea class="richtext-code" v-model="code"></textarea>
			<editor-content :editor="editor" />
		</div>
	</div>
</template>

<!-- =====================
	Scripts
	===================== -->

<script>

import {Editor, EditorContent, EditorMenuBar, EditorMenuBubble} from 'tiptap'
import {
	VerbisBlockquote,
	VerbisBold,
	VerbisBulletList,
	VerbisCode,
	VerbisCodeBlock,
	VerbisCodeBlockHighlight,
	VerbisHardBreak,
	VerbisHeading,
	VerbisHorizontalRule,
	VerbisItalic,
	VerbisLink,
	VerbisListItem,
	VerbisOrderedList,
	VerbisStrike,
	VerbisUnderline,
	VerbisColour,
	VerbisTable
} from '../../../extensions/tiptap/index'
import {
	Highlight,
	History,
	Search,
	TrailingNode,
	TableHeader,
	TableCell,
	TableRow,
} from 'tiptap-extensions';

const Chrome = require('vue-color/src/components/Compact.vue').default;

export default {
	name: "FieldRichText",
	components: {
		EditorContent,
		EditorMenuBar,
		EditorMenuBubble,
		'color-picker': Chrome,
	},
	props: {},
	data() {
		return {
			html: '',
			editor: false,
			config: {},
			linkMenuIsActive: false,
			showTextColorPicker: false,
			palette: false,
			textColor: '#000000',
			code: '',
			codeView: false,
		}
	},
	mounted() {
		this.config = this.getEditorConfig
		this.setUpEditor()
		this.setColourPalette()

	},
	computed: {
		getEditorConfig() {
			return this.$store.state.theme.editor
		}
	},
	methods: {
		setContent(content) {
			if (this.editor) this.editor.setContent(content);
		},
		// getEditorConfig() {
		// 	return new Promise((resolve, reject) => {
		// 		this.axios.get('/theme')
		// 			.then(res => {
		// 				this.config = res.data.data.editor
		// 				resolve()
		// 			})
		// 			.catch(e => {
		// 				console.log(e);
		// 				reject();
		// 			})
		// 	});
		// },
		setUpEditor() {
			const extensions = this.createExtensions();
			this.editor = new Editor({
				content: '',
				onUpdate: ({ getHTML }) => {
					this.html = getHTML()
					this.$emit('update', {body: getHTML()})
				},
				extensions: extensions,
			});
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
			} else {
				this.code = this.html

			}
			this.codeView = !this.codeView
		},
		createExtensions() {
			let extensions = [],
				headingArr = [];

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
		flex-wrap: wrap;
		width: 100%;
		background-color: $white;
		border-top-left-radius: $richtext-border-radius;
		border-top-right-radius: $richtext-border-radius;
		border: 2px solid $grey-light;
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
		border: 2px solid $grey-light;
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
		border: 2px solid $grey-light;
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
		border: 2px solid $grey-light;
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

	&-code {
		position: absolute;
		top: 0;
		left: 0;
		width: 100% !important;
		height: 100% !important;
		border: none;
		padding: 20px;
		opacity: 0;
		z-index: -1;
		color: $copy-light;
		font-size: 14px;
		resize: none;
	}

	// Content (Body)
	// =========================================================================

	&-content {
		position: relative;
		background-color: $white;
		border: 2px solid $grey-light;
		border-bottom-left-radius: $richtext-border-radius;
		border-bottom-right-radius: $richtext-border-radius;
		overflow-y: scroll;

		.ProseMirror {
			min-height: 350px;
			max-height: 700px;
			padding: 20px;
		}

		&-codeview {

			#{$self}-code {
				opacity: 1;
				z-index: 99;
			}
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

		ul {
			display: flex;
			align-items: center;
			list-style: none;

			li {
				line-height: 1;
			}

			&:before {
				position: relative;
				content: "\2022";
				color: $secondary;
				font-weight: bold;
				display: inline-block;
				width: 1em;
				margin-right: 5px;
			}

			p {
				margin-bottom: 0;
			}

		}

		ol {
			margin: 0;
			padding: 0;
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