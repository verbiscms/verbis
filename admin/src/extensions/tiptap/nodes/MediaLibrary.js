import { Node } from 'tiptap'

export default class VerbisMediaLibrary extends Node {

    get name() {
        return 'medialibrary'
    }

    get schema() {
        return {
            // here you have to specify all values that can be stored in this node
            attrs: {
                src: {
                    default: null,
                },
            },
            group: 'block',
            selectable: false,
            // parseDOM and toDOM is still required to make copy and paste work
            parseDOM: [{
                tag: 'iframe',
                getAttrs: dom => ({
                    src: dom.getAttribute('src'),
                }),
            }],
            toDOM: node => ['iframe', {
                src: node.attrs.src,
                frameborder: 0,
                allowfullscreen: 'true',
            }],
        }
    }

    // return a vue component
    // this can be an object or an imported component
    get view() {
        return {
            // there are some props available
            // `node` is a Prosemirror Node Object
            // `updateAttrs` is a function to update attributes defined in `schema`
            // `view` is the ProseMirror view instance
            // `options` is an array of your extension options
            // `selected` is a boolean which is true when selected
            // `editor` is a reference to the TipTap editor instance
            // `getPos` is a function to retrieve the start position of the node
            // `decorations` is an array of decorations around the node
            props: ['node', 'updateAttrs', 'view'],
            computed: {
                src: {
                    get() {
                        return this.node.attrs.src
                    },
                    set(src) {
                        // we cannot update `src` itself because `this.node.attrs` is immutable
                        this.updateAttrs({
                            src,
                        })
                    },
                },
            },
            template: `
        <div class="iframe">
          <iframe class="iframe__embed" :src="src"></iframe>
          <input class="iframe__input" type="text" v-model="src" v-if="view.editable" />
        </div>
      `,
        }
    }

}
