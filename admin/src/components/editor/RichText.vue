<template>
	<div class="editor" v-if="editor">
		<editor-menu-bar :editor="editor" v-slot="{ commands, isActive, focused }">
			<div class="menubar text-nowrap" :class="{ 'is-focused fadeInDownXs': focused }">
				<button class="menubar__button" :class="{ 'is-active': isActive.bold() }" @click="commands.bold">
					<i class="fal fa-bold"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.italic() }" @click="commands.italic">
					<i class="fal fa-italic"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.strike() }" @click="commands.strike">
					<i class="fal fa-strikethrough"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.underline() }" @click="commands.underline">
					<i class="fal fa-underline"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.code() }" @click="commands.code">
					<i class="fal fa-code"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.paragraph() }" @click="commands.paragraph">
					<i class="fal fa-paragraph"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.heading({ level: 1 }) }" @click="commands.heading({ level: 1 })">
					H1
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.heading({ level: 2 }) }" @click="commands.heading({ level: 2 })">
					H2
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.heading({ level: 3 }) }" @click="commands.heading({ level: 3 })">
					H3
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.heading({ level: 4 }) }" @click="commands.heading({ level: 4 })">
					H4
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.bullet_list() }" @click="commands.bullet_list">
					<i class="fal fa-list-ul"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.ordered_list() }" @click="commands.ordered_list">
					<i class="fal fa-list-ol"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.blockquote() }" @click="commands.blockquote">
					<i class="fal fa-quote-right"></i>
				</button>
				<button class="menubar__button" :class="{ 'is-active': isActive.code_block() }" @click="commands.code_block">
					<i class="fal fa-code"></i>
				</button>
				<button class="menubar__button" @click="commands.undo">
					<i class="fal fa-undo"></i>
				</button>
				<button class="menubar__button" @click="commands.redo">
					<i class="fal fa-redo"></i>
				</button>
			</div><!-- end menu bar -->
		</editor-menu-bar>
		<editor-content class="editor__content" :editor="editor"/>
	</div>
</template>

<script>
import { Editor, EditorContent, EditorMenuBar, } from 'tiptap'
import {
	Blockquote,
	BulletList,
	CodeBlock,
	HardBreak,
	Heading,
	ListItem,
	OrderedList,
	TodoItem,
	TodoList,
	Bold,
	Code,
	Italic,
	Link,
	Strike,
	Underline,
	History,
	Placeholder,
} from 'tiptap-extensions'
export default {
	name: "RichText",
	components: {
		EditorContent,
		EditorMenuBar,
	},
	// watch: {
	//     textContent: function() {
	//         alert('You changed something');
	//     },
	// },
	props: {
		key_name: String,
		placeholder: String,
		default: String,
		textContent: String,
	},
	data() {
		return {
			html: 'jj',
			editor: false,
		}
	},
	mounted: function () {
		this.editor = new Editor({
			extensions: [
				new Blockquote(),
				new BulletList(),
				new CodeBlock(),
				new HardBreak(),
				new Heading({ levels: [1, 2, 3, 4, 5] }),
				new ListItem(),
				new OrderedList(),
				new TodoItem(),
				new TodoList(),
				new Link(),
				new Bold(),
				new Code(),
				new Italic(),
				new Strike(),
				new Underline(),
				new History(),
				new Placeholder({
					emptyEditorClass: 'is-editor-empty',
					emptyNodeClass: 'is-empty',
					// emptyNodeText: (this.textContent) ? this.textContent : 'Click here to edit..',
					showOnlyWhenEditable: true,
					showOnlyCurrent: true,
				}),
			],
			onUpdate: ({ getJSON, getHTML }) => {
				this.html = getHTML()
				console.log(getJSON)
			//	this.$emit('text-update', {body: getHTML(), json: getJSON(), key: this.key_name})
			},
			content: (this.textContent) ? this.textContent : this.default,
		})
	},
	methods: {
		setContent(text){
			this.editor.setContent(text)
		},
	},
	beforeUnmount() {
		this.editor.destroy()
	},
}
</script>

<style lang="scss">

</style>