package res

import (
	"github.com/spf13/cobra"
)

var (
	postsCmd = &cobra.Command{
		Use:   "posts",
		Short: "Access the posts",
		Long: ``,
	}
)

// Add child commands
func init() {
	postsCmd.AddCommand(listPostsCmd)
}

// List all the posts within the database
var listPostsCmd = &cobra.Command{
	Use:   "list",
	Short: "List's all users in the database.",
	Run: func(cmd *cobra.Command, args []string) {

		//posts, err := app.store.Posts.GetAll()
		//if err != nil {
		//	log.Error(err)
		//}
		//
		//t := table.NewWriter()
		//t.SetOutputMirror(os.Stdout)
		//t.AppendHeader(table.Row{"#", "Slug", "Title", "Resource", "Page Template", "Layout", "Fields", "Status", "Page Views", "User ID", "Created At", "Updated At"})
		//
		//for _, v := range posts {
		//	t.AppendRows([]table.Row{
		//		{v.Id, v.Slug, v.Title, v.Resource, v.PageTemplate, v.Layout, v.Fields, v.Status, v.UserId, v.CreatedAt, v.UpdatedAt},
		//	})
		//}
		//
		//t.Render()
	},
}
