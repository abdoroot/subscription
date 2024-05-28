// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package customer

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/abdoroot/subscription/views/components"
	"github.com/abdoroot/subscription/views/layouts"
)

func Show() templ.Component {
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
		templ_7745c5c3_Err = ShowHeader().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func FullShowTempl() templ.Component {
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
		templ_7745c5c3_Var3 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = Show().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.AdminApp().Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ShowHeader() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center\"><div class=\"flex-none text-xl\">Ahmed Mohamed Saeed</div><div class=\"grow h-14\"></div><div class=\"flex-none flex space-x-2\"><select class=\"flex items-center gap-2 text-sm bg-secondary rounded bg-primary text-center px-2.5 py-1.5 text-white\"><option>New Transaction</option> <option>All</option> <option>InActive</option> <option>Non-Subscribers</option> <option>Invoiced</option></select>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Button("Edit", components.ColorGray, false, templ.Attributes{"data-target-modal": "create-product-modal"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"text-sm font-semibold text-center border-b  border-secondarygray\"><ul class=\"flex flex-wrap tabs\"><li><a class=\"inline-block py-4 active hover:cursor-pointer\" hx-get=\"/customer/1/overview\" hx-target=\"#tabscontent\">Overview</a></li><li><a class=\"inline-block p-4 hover:cursor-pointer\" hx-get=\"/customer/1/transaction\" hx-target=\"#tabscontent\">Transactions</a></li><li><a class=\"inline-block p-4  hover:cursor-pointer\" hx-get=\"/customer/1/mail\" hx-target=\"#tabscontent\">Mails</a></li><li class=\"me-3\"><a class=\"inline-block p-4\" hx-get=\"/customer/1/statement\" hx-target=\"#tabscontent\">Statement</a></li></ul></div><div class=\"py-4 tabs\" id=\"tabscontent\" hx-target=\"#tabscontent\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = OverviewTab().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func OverviewTab() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex\"><div class=\"w-72 h-auto bg-secondarygray rounded p-2 flex flex-col\"><div class=\"flex text-sm\"><div class=\"header flex-none\"><img class=\"w-10 rounded\" src=\"/static/img/defaultuser.jpeg\" alt=\"default user image\"></div><div class=\"details w-full ml-2 flex flex-col\"><div class=\"font-semibold\">Mr. abdelhadi mohamed</div><div>abd.200930@gmail.com</div><div class=\"flex space-x-1 items-center\"><span><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 430 512\" class=\"w-3\"><path d=\"M310.32 512c-11.94 0-23.92-2.89-34.96-8.72-50.05-26.48-122.17-75.42-181.62-158.16C32.78 260.26 9.58 173.94.86 116.47c-4.25-28.08 7.51-55.75 30.69-72.22l39.7-28.2A86.57 86.57 0 0194.34 4.36c17.31-5.73 36.45-5.81 53.89-.23 18.06 5.78 33.42 17.02 44.41 32.52l24.67 34.77c13.88 19.56 14.24 46.08.9 66-3.89 5.84-8.83 10.92-14.65 15.05-.18.13-.36.25-.55.38l-38.87 25.62c-9.79 6.45-13.17 19.23-7.88 29.74a485.137 485.137 0 0039.12 64.51A484.167 484.167 0 00238.81 325c8.01 8.37 20.64 9.6 30.04 2.93l39.34-27.94c12.66-9 28.06-12.52 43.38-9.93 15.32 2.6 28.71 11.01 37.7 23.68l24.67 34.75c10.99 15.49 16.54 33.7 16.04 52.65-.48 18.3-6.86 36.34-17.96 50.8-.13.17-.27.34-.41.51a87.105 87.105 0 01-18.26 17.44l-39.7 28.2c-12.99 9.23-28.12 13.91-43.33 13.91zM121.79 37.98c-5.15 0-10.37.82-15.48 2.51a49.056 49.056 0 00-12.99 6.58l-39.7 28.2c-11.4 8.1-17.19 21.7-15.1 35.5 8.1 53.34 29.62 133.44 86.16 212.16 55.13 76.74 122.06 122.14 168.5 146.72 12.29 6.5 27.01 5.52 38.4-2.57l39.7-28.2c4.01-2.85 7.54-6.25 10.5-10.1.1-.13.21-.26.31-.39 12.9-17.14 13.23-40.34.78-57.88l-24.67-34.76c-3.1-4.38-7.72-7.28-12.99-8.17-5.27-.89-10.57.32-14.93 3.42l-39.34 27.94a61.12 61.12 0 01-41.26 11 61.395 61.395 0 01-38.37-18.64 524.632 524.632 0 01-46.84-56.39 524.192 524.192 0 01-42.19-69.58c-13.99-27.75-4.99-61.55 20.92-78.63l38.55-25.4a19.85 19.85 0 004.81-5c4.63-6.92 4.51-16.09-.29-22.85L161.6 58.68C152.19 45.42 137.32 38 121.81 38z\"></path></svg></span><span>055xxxxxxxxx</span></div><div class=\"flex space-x-1 items-center\"><span><svg version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" x=\"0\" y=\"0\" viewBox=\"0 0 512 512\" xml:space=\"preserve\" class=\"w-3\"><path d=\"M384 0H160c-17.7 0-32 14.3-32 32v448c0 17.7 14.3 32 32 32h224c17.7 0 32-14.3 32-32V32c0-17.7-14.3-32-32-32zM224 24h96c4.4 0 8 3.6 8 8s-3.6 8-8 8h-96c-4.4 0-8-3.6-8-8s3.6-8 8-8zm48 466c-14.3 0-26-11.7-26-26s11.7-26 26-26 26 11.7 26 26-11.7 26-26 26zm112-90c0 8.8-7.2 16-16 16H176c-8.8 0-16-7.2-16-16V80c0-8.8 7.2-16 16-16h192c8.8 0 16 7.2 16 16v320z\"></path></svg></span><span>055xxxxxxxxx</span></div><div><ul class=\"text-secondaryblue inline-flex items-baseline space-x-1\"><li>Edit</li><li>Send Email</li><li>Delete</li></ul></div></div></div><div><ul><li class=\"mb-3 border-b border-gray\"><button type=\"button\" class=\"flex items-center w-full py-3 text-sm rounded\" aria-controls=\"address_ul\" data-collapse-toggle=\"address_ul\"><span class=\"w-full flex items-center\"><span class=\"flex-none uppercase \">Address</span></span> <svg id=\"Layer_2\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 374.98 227.51\" class=\"flex-none fill-current w-3 mr-2 text-secondaryblue\"><path d=\"M187.46 227.51c-10.23 0-20.46-3.9-28.27-11.71L11.73 68.45C-3.9 52.83-3.91 27.51 11.71 11.88c15.62-15.63 40.94-15.64 56.57-.02l119.18 119.09L306.69 11.72c15.62-15.62 40.95-15.62 56.57 0 15.62 15.62 15.62 40.95 0 56.57L215.75 215.8c-7.81 7.81-18.05 11.72-28.28 11.72z\" id=\"Layer_1-2\"></path></svg></button><ul id=\"address_ul\" class=\"slide-down text-sm\"><li><div class=\"w-full py-2 \">Billing Address</div><div class=\"text-muted\">No Billing Address </div></li><li class=\"pb-3\"><div class=\"w-full py-2\">Shipping Address</div><div class=\"text-muted\">No Shipping Address </div></li></ul></li><li class=\"mb-3 border-b border-gray\"><button type=\"button\" class=\"flex items-center w-full py-3 text-sm rounded\" aria-controls=\"record_ul\" data-collapse-toggle=\"record_ul\"><span class=\"w-full flex items-center\"><span class=\"flex-none uppercase\">Record Info</span></span> <svg id=\"Layer_2\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 374.98 227.51\" class=\"flex-none fill-current w-3 mr-2 text-secondaryblue\"><path d=\"M187.46 227.51c-10.23 0-20.46-3.9-28.27-11.71L11.73 68.45C-3.9 52.83-3.91 27.51 11.71 11.88c15.62-15.63 40.94-15.64 56.57-.02l119.18 119.09L306.69 11.72c15.62-15.62 40.95-15.62 56.57 0 15.62 15.62 15.62 40.95 0 56.57L215.75 215.8c-7.81 7.81-18.05 11.72-28.28 11.72z\" id=\"Layer_1-2\"></path></svg></button><ul id=\"record_ul\" class=\"slide-up text-sm flex flex-row\"><div class=\"flex-1\"><div class=\"text-muted\">Customer ID</div><div class=\"text-muted\">Created on</div><div class=\"text-muted\">Created By</div></div><div class=\"flex-1 pb-3\"><div>12548798</div><div>20 May 2024</div><div>abdelhadi</div></div></ul></li></ul></div></div><div class=\"flex-1 p-2\"><div class=\"font-semibold\">Receivables</div><div class=\"flex flex-col text-xs\"><div class=\"flex flex-row mt-3\"><div class=\"basis-1/3 text-muted uppercase bg-secondarygray border-b border-t border-gray p-2 font-semibold\">Currency</div><div class=\"basis-1/3 text-muted uppercase bg-secondarygray border-b border-t border-gray p-2 font-semibold\">Outstanding Receivables</div></div><div class=\"flex flex-row\"><div class=\"basis-1/3 text-muted uppercase border-b border-gray p-2 font-semibold\">AED- UAE Dirham</div><div class=\"basis-1/3 uppercase  border-b border-t border-gray p-2 font-semibold text-secondaryblue\">AED150</div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func TransactionTab(r string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>Transaction html ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(r)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/customer/show.templ`, Line: 174, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MailTab(r string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>mails html ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(r)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/customer/show.templ`, Line: 178, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func StatementTab(r string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>Statement html ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(r)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/customer/show.templ`, Line: 182, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
