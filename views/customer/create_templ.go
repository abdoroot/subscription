// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package customer

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/abdoroot/subscription/views/layouts"
)

func Create() templ.Component {
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
		templ_7745c5c3_Err = CreateHeader().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col text-sm\"><div class=\"flex pt-6\"><div class=\"w-2/12\">Customer Type</div><div class=\"w-4/12\"><div class=\"flex space-x-2\"><div class=\"flex items-center\"><input checked type=\"radio\" value=\"business\" name=\"customer-type\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-3\"> <label class=\"ms-2\">Business</label></div><div class=\"flex items-center\"><input type=\"radio\" value=\"individual\" name=\"customer-type\" class=\"w-4 h-4 focus:bg-secondaryblue focus:ring-2\"> <label class=\"ms-2\">Individual</label></div></div></div></div><div class=\"flex pt-4 items-center\"><div class=\"w-2/12\">Primary Contact </div><div class=\"w-4/12\"><div class=\"flex space-x-2\"><div class=\"\"><select name=\"\" class=\"bg-white border rounded border:secondarygray py-1.5 px-2\"><option>Mr.</option> <option>Mrs.</option> <option>Miss.</option> <option>Dr.</option></select></div><div class=\"\"><input type=\"text\" name=\"first-name\" placeholder=\"First name\" class=\"rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div><div class=\"\"><input type=\"text\" name=\"last-name\" placeholder=\"Last name\" class=\"rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div></div></div><div class=\"flex pt-4 items-center\"><div class=\"w-2/12 text-red-500\">Customer Display Name *</div><div class=\"w-3/12 flex\"><input type=\"text\" name=\"display-name\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div><div class=\"flex pt-4 items-center\"><div class=\"w-2/12 text-red-500\">Customer Email *</div><div class=\"w-3/12 flex\"><input type=\"text\" name=\"email\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div><div class=\"flex pt-4 items-center\"><div class=\"w-2/12\">Customer Phone</div><div class=\"w-3/12 flex space-x-2\"><input type=\"text\" name=\"email\" placeholder=\"Work phone\" class=\"w-6/12 rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"> <input type=\"text\" name=\"email\" placeholder=\"mobile\" class=\"w-6/12 rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div><div class=\"pt-6\"><ul class=\"flex flex-wrap border-b border-gray-200 normaltabs\"><li class=\"mr-2\"><a class=\"inline-block py-4 px-4 cursor-pointer active\" data-tab=\"address-tap\">Address</a></li><li class=\"mr-2\"><a class=\"inline-block py-4 px-4  border-transparent cursor-pointer\" data-tab=\"other-details-tap\">Other Details</a></li></ul><div><div class=\"p-4 rounded-lg tab-content flex\" id=\"address-tap\" role=\"tabpanel\"><div class=\"left w-6/12\"><span class=\"font-bold text-lg\">Billing Address</span><div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">Country / Region</div><div class=\"w-6/12\"><select class=\"w-full bg-white rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\" data-twe-select-filter=\"true\"><option value=\"1\">Option 1</option> <option value=\"2\">Option 2</option></select></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">Address</div><div class=\"w-6/12\"><textarea type=\"text\" name=\"\" placeholder=\"Street 1\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></textarea></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\"></div><div class=\"w-6/12\"><textarea type=\"text\" name=\"\" placeholder=\"Street 2\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></textarea></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">State</div><div class=\"w-6/12\"><input type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">Zip code</div><div class=\"w-6/12\"><input type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">Phone</div><div class=\"w-6/12\"><input type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div></div></div><div class=\"right w-6/12\"><span class=\"font-bold text-lg\">Shipping Address  (<span class=\" text-secondaryblue text-sm font-medium\"><a>Copy billing address</a></span>)</span><div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">Country / Region</div><div class=\"w-6/12\"><select class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\" data-twe-select-filter=\"true\"><option value=\"1\">Option 1</option> <option value=\"2\">Option 2</option></select></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">Address</div><div class=\"w-6/12\"><textarea type=\"text\" name=\"\" placeholder=\"Street 1\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></textarea></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\"></div><div class=\"w-6/12\"><textarea type=\"text\" name=\"\" placeholder=\"Street 2\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></textarea></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">State</div><div class=\"w-6/12\"><input type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">Zip code</div><div class=\"w-6/12\"><input type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div><div class=\"flex items-center pt-4\"><div class=\"w-4/12\">Phone</div><div class=\"w-6/12\"><input type=\"text\" name=\"\" class=\"w-full rounded border border:secondarygray py-1.5 px-2 focus:border-secondaryblue\"></div></div></div></div></div><div class=\"hidden p-4 rounded-lg tab-content\" id=\"other-details-tap\" role=\"tabpanel\">Other Details-This is some placeholder content for the Mails tab. Clicking another tab will toggle the visibility of this one for the next. The tab JavaScript swaps classes to control the content visibility and styling.</div></div><br><br><br><br></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CreateFooter().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CreateHeader() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"border-b border:secondarygray py-2\"><div class=\"text-2xl\">New Customer</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CreateFooter() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"fixed bottom-0 px-3 righ-0 w-full h-12 bg-white shadow-md flex items-center text-sm border-t border-gray\"><button class=\" bg-secondary rounded  px-2.5 py-1.5 text-white\">Save</button> <button class=\"bg-secondary rounded  px-2.5 py-1.5 text-white ml-2\">Save and Subscribe</button> <button class=\"bg-secondarygray rounded border border-gray px-2.5 py-1.5 ml-2\">Cancel</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func FullCreateTempl() templ.Component {
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
		templ_7745c5c3_Var5 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = Create().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.AdminApp().Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
