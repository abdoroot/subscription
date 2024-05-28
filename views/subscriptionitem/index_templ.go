// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package subscriptionitem

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/abdoroot/subscription/views/components"
	"github.com/abdoroot/subscription/views/layouts"
)

func Index() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"page\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CreateModal().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex\"><div class=\"w-2/6 sticky top-0 h-screen border-r border-secondarygraydarker\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ItemsHeader().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col\"><div class=\"flex flex-col items-start bg-secondarygray hover:bg-secondarygray py-6 px-2 border-b border-secondarygraydarker hover:cursor-pointer\"><span class=\"font-medium\">Vpn Service</span> <span class=\"text-sm w-auto mt-2 bg-secondarygray border border-secondarygraydarker p-0.5 hover:bg-secondaryblue  hover:text-white\">1 plan</span></div><div class=\"flex flex-col items-start py-6 px-2 border-b border-secondarygraydarker hover:bg-secondarygray hover:cursor-pointer\"><span class=\"font-medium\">Premuim Gym</span> <span class=\"text-sm w-auto mt-2 bg-secondarygray border border-secondarygraydarker p-0.5 hover:bg-secondaryblue hover:text-white\">2 plan</span></div><div class=\"flex flex-col items-start py-6 px-2 border-b border-secondarygraydarker hover:bg-secondarygray hover:cursor-pointer\"><span class=\"font-medium\">Silver</span> <span class=\"text-sm w-auto mt-2 bg-secondarygray border border-secondarygraydarker p-0.5 hover:bg-secondaryblue hover:text-white\">2 plan</span></div></div></div><div class=\"w-5/6\" id=\"content-column\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ItemDescHeader().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex py-3 bg-secondarygray border-b border-secondarygraydarker text-sm\"><div class=\"items-center flex flex-col w-2/12\"><span class=\"text-2xl\">2</span> <span>Plans</span></div><div class=\"items-center flex flex-col w-2/12 text-secondaryblue\"><span class=\"text-2xl\">--</span> <span>Addons</span></div><div class=\"items-center flex flex-col w-2/12 text-secondaryblue\"><span class=\"text-2xl\">--</span> <span>Coupons</span></div></div><div class=\"pt-10 px-6\"><span class=\"uppercase text-xl\">plans</span><div class=\"mt-2 bg-secondarygray border border-secondarygraydarker p-3 flex\"><span class=\"w-4/6 text-muted\">Plan Details</span> <span class=\"w-2/6 text-muted text-right\">Price</span></div><div class=\"p-3 flex border-b border-secondarygraydarker\"><div class=\"w-4/6 flex items-center\"><div class=\"w-2/12\"><svg version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" x=\"0\" y=\"0\" viewBox=\"0 0 512 512\" xml:space=\"preserve\" class=\"w-16 p-2 bg-secondarygray\"><path d=\"M417.2 378.8H95.3c-7 0-12.8-5.7-12.8-12.8v-34.9c0-2.7.8-5.2 2.3-7.4l44.6-63c4-5.6 11.6-7 17.4-3.3l60.8 39.7c4.9 3.2 11.1 2.7 15.5-1.1l116.8-103.2c5.5-4.9 14.1-4.1 18.5 1.8l66.3 86c1.7 2.2 2.7 5 2.7 7.8v80.2c0 5.6-4.6 10.2-10.2 10.2z\" fill=\"#40bab5\"></path><path d=\"M212.2 157.7c23.2 0 42 19 42 42.4s-18.8 42.4-42 42.4-42-19-42-42.4c.1-23.4 18.9-42.4 42-42.4z\" fill=\"#fbbe01\"></path><path d=\"M462 60.8c16.5 0 30 13.5 30 30V422c0 16.5-13.5 30-30 30H50.4c-16.5 0-30-13.5-30-30V90.8c0-16.5 13.5-30 30-30H462m0-20H50.4c-27.6 0-50 22.4-50 50V422c0 27.6 22.4 50 50 50H462c27.6 0 50-22.4 50-50V90.8c0-27.6-22.4-50-50-50z\" fill=\"#888\"></path></svg></div><div class=\"w-10/12 flex flex-col\"><span class=\"text-base font-medium\">monthly</span> <span class=\"text-muted\">Plan code: m</span><ul class=\"text-secondaryblue inline-flex items-baseline space-x-1 text-sm\"><li>Edit</li><li>Mark as InActive</li><li>Share</li><li>Delete</li></ul></div></div><div class=\"w-2/6 text-right flex flex-col\"><span class=\"font-medium\">AED100</span> <span class=\"text-muted text-sm\">per month</span></div></div></div><div class=\"pt-10 px-6\"><span class=\"uppercase text-xl\">ADDONS</span><div class=\"mt-2 bg-secondarygray border border-secondarygraydarker p-3 flex\"><span class=\"w-4/6 text-muted\">Addon Details</span> <span class=\"w-2/6 text-muted text-right\">Price</span></div><div class=\"p-3 flex\"><div class=\"w-full text-center\">No addons yet!</div></div></div><div class=\"pt-10 px-6\"><span class=\"uppercase text-xl\">COUPONS</span><div class=\"mt-2 bg-secondarygray border border-secondarygraydarker p-3 flex\"><span class=\"w-4/6 text-muted\">Addon Details</span> <span class=\"w-2/6 text-muted text-right\">Price</span></div><div class=\"p-3 flex\"><div class=\"w-full text-center\">No coupons yet!</div></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ItemDescHeader() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center border-b-2 border-secondarygraydarker px-2 sticky top-0\"><div class=\"flex-none\"><div class=\"text-xl\">Vpn Service</div></div><div class=\"grow h-14\"></div><div class=\"flex space-x-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Button("Edit", components.ColorGray, false, templ.Attributes{"data-target-modal": "create-product-modal"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Button("Add plan", components.ColorGreen, false, templ.Attributes{
			"hx-get":       "/subscription-item/plan/create",
			"hx-target":    "#content-column",
			"hx-indicator": "#spinner",
			"hx-push-url":  "true",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Button("Add addon", components.ColorGray, false, nil).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ItemsHeader() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center border-b-2 border-secondarygraydarker px-2\"><div class=\"flex-none\"><select class=\"bg-white py-2 rounded\"><option>All Products</option> <option>Active</option> <option>InActive</option></select></div><div class=\"grow h-14\"></div><div class=\"flex-none\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Button("New", components.ColorGreen, true, templ.Attributes{"data-target-modal": "create-product-modal"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- <button data-target-modal=\"create-product-modal\"\n\t\t\tclass=\"open-modal flex items-center gap-2 text-sm bg-secondary rounded px-2.5 py-1.5 text-white\">\n\t\t\t<svg class=\"fill-current w-3 h-3\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\n\t\t\t\t<path\n\t\t\t\t\td=\"M15 7H9V1C9 0.4 8.6 0 8 0C7.4 0 7 0.4 7 1V7H1C0.4 7 0 7.4 0 8C0 8.6 0.4 9 1 9H7V15C7 15.6 7.4 16 8 16C8.6 16 9 15.6 9 15V9H15C15.6 9 16 8.6 16 8C16 7.4 15.6 7 15 7Z\"\n\t\t\t\t\tfill=\"\"></path>\n\t\t\t</svg>\n\t\t\tNew\n\t\t</button> --></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CreateModal() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"create-product-modal\" tabindex=\"-1\" class=\"hidden modal overflow-y-auto overflow-x-hidden fixed top-1/4 left-1/3 z-50 justify-center items-center w-full h-[calc(100%-1rem)] max-h-full\"><div class=\"relative p-4 w-full max-w-md max-h-full\"><div class=\"relative bg-white rounded-lg shadow\"><div class=\"bg-secondarygray border-b border-secondarygraydarker py-3 px-2\"><span class=\"text-lg font-medium\">New Product</span> <button type=\"button\" class=\"close-modal absolute top-2 end-2.5 bg-transparent  rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center\" data-modal-hide=\"popup-modal\"><svg class=\"w-3 h-3 text-red-500\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 14 14\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6\"></path></svg> <span class=\"sr-only\">Close modal</span></button></div><div class=\"flex flex-col text-sm w-full py-2 px-3\"><div class=\"pt-4\"><div class=\"text-red-500\">Name *</div><div class=\"\"><input type=\"text\" class=\"py-1.5 w-full rounded border border:black focus:border-secondaryblue\"></div></div><div class=\" pt-4\"><div class=\"\">Description</div><div class=\"\"><textarea type=\"text\" name=\"\" placeholder=\"Description\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></textarea></div></div><div class=\"mt-2 py-4 border-t border-secondarygraydarker\"><button class=\" bg-secondary rounded  px-2.5 py-1.5 text-white\">Save</button> <button class=\"bg-secondarygray rounded border border-gray px-2.5 py-1.5 ml-2 close-modal\">Cancel</button></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func FullTempl() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var6 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = Index().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.AdminApp().Render(templ.WithChildren(ctx, templ_7745c5c3_Var6), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
