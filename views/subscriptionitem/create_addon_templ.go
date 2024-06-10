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

func CreateAddon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"page\" class=\"px-3 py-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CreateAddonHeader().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col text-sm pb-5\"><div class=\"flex pt-6 items-center\"><div class=\"w-2/12 text-red-500\">Addon Name*</div><div class=\"w-4/12 flex\"><input type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div><div class=\"flex pt-4 items-center\"><div class=\"w-2/12 text-red-500\">Addon Code*</div><div class=\"w-4/12 flex\"><input type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div><div class=\"w-6/12 text-muted px-4\">A unique code of your choice which lets you identify this addon.</div></div><div class=\"flex pt-4 items-center\"><div class=\"w-2/12 text-red-500\">Type*</div><div class=\"w-4/12 flex\"><div class=\"flex space-x-2\"><div class=\"flex items-center\"><input checked type=\"radio\" value=\"business\" name=\"plan_type\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-3\"> <label class=\"ms-2\">Goods</label></div><div class=\"flex items-center\"><input type=\"radio\" value=\"individual\" name=\"plan_type\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-2\"> <label class=\"ms-2\">Service</label></div></div></div></div><div class=\"flex pt-4 items-center\"><div class=\"w-2/12\">Addon Description</div><div class=\"w-4/12 flex\"><textarea type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-4 px-2 focus:border-secondaryblue\"></textarea></div></div><hr class=\"h-auto my-8 w-full border-b border-secondarygraydarker\"><div class=\"flex uppercase text-base font-medium text-muted\"><span>PRICING DETAILS</span></div><div class=\"flex pt-6 items-center\"><div class=\"w-2/12\">Pricing Model</div><div class=\"w-4/12 flex\"><select class=\"w-full bg-white rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\" data-twe-select-filter=\"true\"><option value=\"1\">Option 1</option> <option value=\"2\">Option 2</option></select></div></div><div class=\"flex pt-4 items-center\"><div class=\"w-2/12  text-red-500\">Price*</div><div class=\"w-4/12 flex items-center\"><div class=\"bg-secondarygraydarker border-r border-t border-b border:secondarygraydarker px-2 py-1.5 rounded-l\">AED</div><input type=\"text\" class=\"py-1.5 px-2 w-full rounded-r border-r border-t border-b border:secondarygray  focus:border-secondaryblue\"></div><div class=\"w-6/12 text-muted px-4\">/unit/month</div></div><div class=\"flex pt-6 items-center\"><div class=\"w-2/12 text-red-500\">Unit name*</div><div class=\"w-4/12 flex\"><select class=\"w-full bg-white rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\" data-twe-select-filter=\"true\"><option value=\"1\">Option 1</option> <option value=\"2\">Option 2</option></select></div><div class=\"w-6/12 text-muted px-4\">A name of your choice to refer to one unit of the addon. Eg: Stamp is a unit for snail mail add-on.</div></div><hr class=\"h-auto my-8 w-full border-b border-secondarygraydarker\"><div class=\"flex uppercase text-base font-medium text-muted\"><span>BILLING PREFERENCES</span></div><div class=\"flex pt-6 items-center\"><div class=\"w-2/12  text-red-500\">Addon Type*</div><div class=\"w-4/12 flex\"><div class=\"flex space-x-2\"><div class=\"flex items-center\"><input checked type=\"radio\" value=\"business\" name=\"addon_type\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-3\"> <label class=\"ms-2\">One-time</label></div><div class=\"flex items-center\"><input type=\"radio\" value=\"individual\" name=\"addon_type\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-2\"> <label class=\"ms-2\">Recurring</label></div></div></div></div><div class=\"flex pt-4\"><div class=\"w-2/12  text-red-500\">Pricing Interval*</div><div class=\"w-4/12 flex items-center space-x-2\"><select class=\"w-full bg-white rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\" data-twe-select-filter=\"true\"><option value=\"1\">Week(s)</option> <option value=\"1\" selected>Month(s)</option> <option value=\"2\">Year 2</option></select></div></div><div class=\"flex pt-4\"><div class=\"w-2/12\"><span>Associated Plans*</span></div><div class=\"w-4/12 flex\"><div class=\"flex flex-col space-y-2\"><div class=\"flex\"><input checked type=\"checkbox\" value=\"\" name=\"\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-3\"> <label class=\"ms-2\">All Plans</label></div><div id=\"other_plans mt-6\"><div class=\"flex \"><input checked type=\"checkbox\" value=\"\" name=\"\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-3\"> <label class=\"ms-2\">monthly Plan</label></div><div class=\"flex\"><input checked type=\"checkbox\" value=\"\" name=\"\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-2\"> <label class=\"ms-2\">Yearly plan</label></div></div></div></div></div><hr class=\"h-auto my-6 w-full border-b border-secondarygraydarker\"><div class=\"flex items-center\"><div class=\"w-2/12\"></div><div class=\"w-4/12 flex space-x-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Button("Save", components.ColorGreen, false, nil).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Button("Cancel", components.ColorGray, false, nil).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CreateAddonHeader() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"border-b border:secondarygray py-2\"><div class=\"text-2xl font-semibold\">Add Addon</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func FullCreateAddon() templ.Component {
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
		templ_7745c5c3_Var4 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = CreateAddon().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.AdminApp().Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
