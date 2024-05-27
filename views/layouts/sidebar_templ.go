// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SideBar() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Aside Menu --><aside class=\"w-1/6 sticky top-0 border-r border-gray-100 bg-primary text-white overflow-y-auto\"><div class=\"py-3 px-6 bg-primarydarker\">Logo</div><div class=\"p-2\"><ul><li class=\"mb-3 rounded transition hover:bg-primarydarker\"><a href=\"/\" class=\"text-sm flex items-center py-2 px-3\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 514 458.5\" class=\"fill-current w-6 h-4 mr-2\" id=\"\"><path d=\"M507 185.7L315.7 20.9C283.6-6.7 236.8-7 204.3 20.2L7.2 185.5c-8.5 7.1-9.6 19.7-2.5 28.2 4 4.7 9.6 7.2 15.3 7.2s9.1-1.5 12.8-4.7l25.6-21.4v165.4c.9 54.4 45.4 98.4 100 98.4h197.2c54.6 0 99.1-44 100-98.4V194.3l25.4 21.8c8.4 7.2 21 6.3 28.2-2.1 7.2-8.4 6.3-21-2.1-28.2zM298.7 418.5h-83.5v-65.7c0-23 18.7-41.8 41.8-41.8s41.8 18.7 41.8 41.8v65.7zm116.8-60c0 33.1-26.9 60-60 60h-16.8v-65.7c0-45.1-36.7-81.8-81.8-81.8s-81.8 36.7-81.8 81.8v65.7h-16.8c-33.1 0-60-26.9-60-60V161.2L230 50.9c17.4-14.6 42.4-14.4 59.6.4l126 108.5v198.7z\" id=\"Layer_1-2\"></path></svg> Home</a></li><li class=\"mb-3 flex rounded hover:bg-primarydarker items-center\"><a href=\"#\" hx-get=\"/customer\" hx-target=\"#main_content\" hx-indicator=\"#spinner\" hx-swap=\"innerHTML\" hx-push-url=\"true\" class=\"grow flex items-center py-2 px-3 text-sm \"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 431.2 512\" class=\"fill-current w-6 h-4 mr-2\"><g id=\"Layer_1-2\"><path d=\"M215.6 227.8c-62.8 0-113.9-51.1-113.9-113.9S152.8 0 215.6 0s113.9 51.1 113.9 113.9-51.1 113.9-113.9 113.9zm0-187.8c-40.8 0-73.9 33.2-73.9 73.9s33.2 73.9 73.9 73.9 73.9-33.2 73.9-73.9S256.3 40 215.6 40zM358 512H73.2c-23.1 0-44.3-10.5-58.3-28.9s-18.4-41.7-12.2-63.9l14.5-51.9c7.4-26.6 23-49.5 44.9-66.1s48.2-25.5 75.8-25.5h155.3c27.6 0 53.8 8.8 75.8 25.5s37.5 39.6 44.9 66.1l14.5 51.9c6.2 22.2 1.8 45.5-12.2 63.9S381 512 357.9 512zM138 315.6c-38.1 0-71.9 25.7-82.2 62.4l-14.5 51.9c-2.8 10.1-.8 20.6 5.5 29 6.3 8.3 16 13.1 26.4 13.1h284.7c10.5 0 20.1-4.8 26.4-13.1 6.3-8.3 8.3-18.9 5.5-29L375.3 378c-10.2-36.7-44-62.4-82.2-62.4H137.8z\"></path><path d=\"M310 421.8h-2c-.6 0-1.3-.2-1.9-.3-.6-.1-1.3-.3-1.9-.5-.6-.2-1.2-.4-1.8-.7-.6-.2-1.2-.5-1.8-.8-.6-.3-1.1-.6-1.7-1-.5-.4-1.1-.8-1.6-1.2-.5-.4-1-.9-1.5-1.3-.5-.5-.9-1-1.3-1.5s-.8-1-1.2-1.6c-.4-.5-.7-1.1-1-1.7-.3-.6-.6-1.2-.8-1.8-.2-.6-.5-1.2-.7-1.8-.2-.6-.4-1.3-.5-1.9-.1-.6-.2-1.3-.3-1.9v-2-2c0-.6.2-1.3.3-1.9.1-.6.3-1.3.5-1.9.2-.6.4-1.2.7-1.8.2-.6.5-1.2.8-1.8.3-.6.6-1.1 1-1.7.4-.5.8-1.1 1.2-1.6.4-.5.9-1 1.3-1.5.5-.5 1-.9 1.5-1.3s1-.8 1.6-1.2c.5-.4 1.1-.7 1.7-1 .6-.3 1.2-.6 1.8-.8.6-.2 1.2-.5 1.8-.7.6-.2 1.3-.4 1.9-.5.6-.1 1.3-.2 1.9-.3 1.3-.1 2.6-.1 3.9 0 .6 0 1.3.2 1.9.3.6.1 1.3.3 1.9.5.6.2 1.2.4 1.8.7.6.2 1.2.5 1.8.8.6.3 1.1.6 1.7 1 .5.4 1.1.8 1.6 1.2.5.4 1 .9 1.5 1.3.5.5.9 1 1.3 1.5s.8 1 1.2 1.6c.4.5.7 1.1 1 1.7.3.6.6 1.2.8 1.8.2.6.5 1.2.7 1.8.2.6.4 1.3.5 1.9.1.6.2 1.3.3 1.9v4c0 .6-.2 1.3-.3 1.9-.1.6-.3 1.3-.5 1.9-.2.6-.4 1.2-.7 1.8-.2.6-.5 1.2-.8 1.8-.3.6-.6 1.1-1 1.7-.4.5-.8 1.1-1.2 1.6-.4.5-.9 1-1.3 1.5-.5.5-1 .9-1.5 1.3s-1 .8-1.6 1.2c-.5.4-1.1.7-1.7 1-.6.3-1.2.6-1.8.8-.6.2-1.2.5-1.8.7-.6.2-1.3.4-1.9.5-.6.1-1.3.2-1.9.3h-2z\"></path></g></svg> Customers</a> <a class=\"flex-none py-2 px-3 opacity-0 hover:opacity-100\" hx-get=\"/customer/create\" hx-target=\"#main_content\" hx-indicator=\"#spinner\" hx-push-url=\"true\" href=\"/customer/create\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" class=\"fill-current w-4 h-4\"><path d=\"M255.8 40.8c-118.9 0-215.4 96.5-215.4 215.4s96.5 215.4 215.4 215.4 215.4-96.4 215.4-215.4S374.7 40.8 255.8 40.8zm94.5 237h-72.9v72.9c-.3 11.9-10.1 21.4-22.1 21.2-11.6-.2-20.9-9.6-21.2-21.2v-72.9h-72.9c-11.9 0-21.6-9.7-21.6-21.6s9.7-21.6 21.6-21.6h72.9v-72.9c.3-11.9 10.1-21.4 22.1-21.2 11.6.2 20.9 9.6 21.2 21.2v72.9h72.9c11.9 0 21.6 9.7 21.6 21.6 0 12-9.7 21.6-21.6 21.6z\"></path></svg></a></li><li class=\"mb-3\"><button type=\"button\" class=\"flex items-center w-full p-2 px-3 text-sm rounded hover:bg-primarydarker\" aria-controls=\"products_items\" data-collapse-toggle=\"products_items\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 504.9 487.3\" class=\"fill-current w-6 h-4 mr-2\"><g id=\"Layer_1-2\"><path d=\"M483 155.4c-19.1-24-47.6-37.7-78.2-37.7H375C373.8 86.8 361.2 58 339.2 36 315.6 12.4 284.3-.4 250.9 0c-32.6.4-63.1 13.6-86 37.1-21.2 21.9-33.5 50.2-35.1 80.5H100c-30.6 0-59.1 13.7-78.2 37.7-19 23.9-26 54.7-19.3 84.5l34.3 169.4c0 .2 0 .3.1.5 5 22 17.5 41.8 35.2 55.8s39.8 21.8 62.3 21.8h236.1c22.5 0 44.7-7.7 62.3-21.8 17.6-14 30.1-33.9 35.2-55.8 0-.2 0-.3.1-.5l34.3-169.4c6.8-29.8-.3-60.5-19.3-84.4zM251.4 40c22.5-.3 43.6 8.3 59.5 24.2 14.4 14.4 22.8 33.3 24.1 53.4H170c3.2-42.9 38.5-77.2 81.5-77.6zm211.9 191.1c0 .2 0 .3-.1.5L428.9 401c-6.4 27.3-30.4 46.3-58.4 46.3H134.4c-28.1 0-52-19-58.4-46.3L41.7 231.6c0-.2 0-.3-.1-.5-4.1-17.9 0-36.4 11.5-50.8 11.4-14.4 28.6-22.6 46.9-22.6h304.8c18.4 0 35.5 8.2 46.9 22.6 11.4 14.4 15.7 32.9 11.5 50.8z\"></path><path d=\"M308.6 282.5H196.3c-11 0-20 9-20 20s9 20 20 20h112.3c11 0 20-9 20-20s-9-20-20-20z\"></path></g></svg> <span class=\"w-full flex items-center\"><span class=\"flex-none  hover:bg-primarydarker\">Products </span></span> <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" class=\"flex-none fill-current w-2 h-2 mr-2 \"><path d=\"M481.6 74.4H30.4c-9.6 0-15.3 10.6-10 18.6l223.3 339.2c4.7 7.2 15.2 7.2 20 .1 39.5-58.8 184.2-274.2 227.8-339.2 5.4-7.9-.3-18.7-9.9-18.7z\"></path></svg></button><ul id=\"products_items\" class=\"slide-up text-sm\"><li class=\"flex rounded hover:bg-primarydarker items-center\"><a hx-get=\"/item\" hx-target=\"#main_content\" hx-indicator=\"#spinner\" hx-swap=\"innerHTML\" hx-push-url=\"true\" class=\"w-full py-2 px-6\">items</a> <a hx-get=\"/item/create\" hx-target=\"#main_content\" hx-indicator=\"#spinner\" hx-push-url=\"true\" class=\"flex-none py-2 px-3 opacity-0 hover:opacity-100\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" class=\"fill-current w-4 h-4\"><path d=\"M255.8 40.8c-118.9 0-215.4 96.5-215.4 215.4s96.5 215.4 215.4 215.4 215.4-96.4 215.4-215.4S374.7 40.8 255.8 40.8zm94.5 237h-72.9v72.9c-.3 11.9-10.1 21.4-22.1 21.2-11.6-.2-20.9-9.6-21.2-21.2v-72.9h-72.9c-11.9 0-21.6-9.7-21.6-21.6s9.7-21.6 21.6-21.6h72.9v-72.9c.3-11.9 10.1-21.4 22.1-21.2 11.6.2 20.9 9.6 21.2 21.2v72.9h72.9c11.9 0 21.6 9.7 21.6 21.6 0 12-9.7 21.6-21.6 21.6z\"></path></svg></a></li><li class=\"flex rounded hover:bg-primarydarker items-center\"><a hx-get=\"/subscription-item\" hx-target=\"#main_content\" hx-indicator=\"#spinner\" hx-swap=\"innerHTML\" hx-push-url=\"true\" class=\"w-full py-2 px-6\">subscription items</a> <a hx-get=\"/subscription-item/create\" hx-target=\"#main_content\" hx-indicator=\"#spinner\" hx-swap=\"innerHTML\" hx-push-url=\"true\" class=\"flex-none  py-2 px-3 opacity-0 hover:opacity-100\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" class=\"fill-current w-4 h-4\"><path d=\"M255.8 40.8c-118.9 0-215.4 96.5-215.4 215.4s96.5 215.4 215.4 215.4 215.4-96.4 215.4-215.4S374.7 40.8 255.8 40.8zm94.5 237h-72.9v72.9c-.3 11.9-10.1 21.4-22.1 21.2-11.6-.2-20.9-9.6-21.2-21.2v-72.9h-72.9c-11.9 0-21.6-9.7-21.6-21.6s9.7-21.6 21.6-21.6h72.9v-72.9c.3-11.9 10.1-21.4 22.1-21.2 11.6.2 20.9 9.6 21.2 21.2v72.9h72.9c11.9 0 21.6 9.7 21.6 21.6 0 12-9.7 21.6-21.6 21.6z\"></path></svg></a></li></ul></li><li class=\"mb-3\"><button type=\"button\" class=\"flex items-center w-full p-2 px-3 text-sm rounded hover:bg-primarydarker\" aria-controls=\"sales_items\" data-collapse-toggle=\"sales_items\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 484.4\" class=\"flex-none fill-current w-4 h-4 mr-2 \"><g id=\"Layer_1-2\"><path d=\"M497.2 95.7c-12.8-15.9-31.9-25.1-52.3-25.1h-202c-11 0-20 9-20 20s9 20 20 20h202c8.3 0 16 3.7 21.2 10.1 5.2 6.4 7.1 14.8 5.3 22.8l-34.5 156c-5.2 23.4-25.5 39.7-49.5 39.7H210.2c-24.4 0-45.7-17.4-50.6-41.3L108 44.4C102.7 18.7 79.9 0 53.6 0H20C9 0 0 9 0 20s9 20 20 20h33.6c7.3 0 13.7 5.2 15.1 12.4l51.7 253.7c8.7 42.5 46.4 73.3 89.8 73.3h177.2c20.6 0 40.7-7.1 56.8-20 16-12.9 27.3-31 31.7-51.1l34.5-156c4.4-20-.4-40.6-13.2-56.5z\"></path><circle cx=\"182.8\" cy=\"450.5\" r=\"33.9\"></circle><path d=\"M414.8 416.6c-18.7 0-33.9 15.2-33.9 33.9s15.2 33.9 33.9 33.9 33.9-15.2 33.9-33.9-15.2-33.9-33.9-33.9z\"></path></g></svg> <span class=\"w-full flex items-center\"><span class=\"flex-none  hover:bg-primarydarker\">Sales </span></span> <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" class=\"flex-none fill-current w-2 h-2 mr-2 \"><path d=\"M481.6 74.4H30.4c-9.6 0-15.3 10.6-10 18.6l223.3 339.2c4.7 7.2 15.2 7.2 20 .1 39.5-58.8 184.2-274.2 227.8-339.2 5.4-7.9-.3-18.7-9.9-18.7z\"></path></svg></button><ul id=\"sales_items\" class=\"slide-up text-sm\"><li class=\"flex rounded hover:bg-primarydarker items-center\"><a href=\"#\" class=\"w-full py-2 px-6\">Invoices</a> <a class=\"flex-none py-2 px-3 opacity-0 hover:opacity-100\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" class=\"fill-current w-4 h-4\"><path d=\"M255.8 40.8c-118.9 0-215.4 96.5-215.4 215.4s96.5 215.4 215.4 215.4 215.4-96.4 215.4-215.4S374.7 40.8 255.8 40.8zm94.5 237h-72.9v72.9c-.3 11.9-10.1 21.4-22.1 21.2-11.6-.2-20.9-9.6-21.2-21.2v-72.9h-72.9c-11.9 0-21.6-9.7-21.6-21.6s9.7-21.6 21.6-21.6h72.9v-72.9c.3-11.9 10.1-21.4 22.1-21.2 11.6.2 20.9 9.6 21.2 21.2v72.9h72.9c11.9 0 21.6 9.7 21.6 21.6 0 12-9.7 21.6-21.6 21.6z\"></path></svg></a></li><li class=\"flex rounded hover:bg-primarydarker items-center\"><a href=\"#\" class=\"w-full py-2 px-6\">Subscriptions</a> <a class=\"flex-none  py-2 px-3 opacity-0 hover:opacity-100\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" class=\"fill-current w-4 h-4\"><path d=\"M255.8 40.8c-118.9 0-215.4 96.5-215.4 215.4s96.5 215.4 215.4 215.4 215.4-96.4 215.4-215.4S374.7 40.8 255.8 40.8zm94.5 237h-72.9v72.9c-.3 11.9-10.1 21.4-22.1 21.2-11.6-.2-20.9-9.6-21.2-21.2v-72.9h-72.9c-11.9 0-21.6-9.7-21.6-21.6s9.7-21.6 21.6-21.6h72.9v-72.9c.3-11.9 10.1-21.4 22.1-21.2 11.6.2 20.9 9.6 21.2 21.2v72.9h72.9c11.9 0 21.6 9.7 21.6 21.6 0 12-9.7 21.6-21.6 21.6z\"></path></svg></a></li></ul></li><li class=\"mb-3 flex rounded hover:bg-primarydarker items-center\"><a href=\"#\" hx-get=\"/customer\" hx-target=\"#main_content\" hx-indicator=\"#spinner\" hx-swap=\"innerHTML\" hx-push-url=\"true\" class=\"grow flex items-center py-2 px-3 text-sm \"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 440.853\" class=\"fill-current w-6 h-4 mr-2\"><g id=\"Layer_1-2\"><path d=\"M466.5 13.763c-78.44-30.13-184.22-3.69-247.42 12.1-15.86 3.96-29.55 7.38-38.34 8.78-22.73 3.6-64.58 3.26-117.83-.94-16.17-1.28-32.28 4.3-44.19 15.3-11.9 11-18.72 26.6-18.72 42.8v268.84c0 29.29 18.28 55.99 45.5 66.45 26.02 9.99 55.04 13.76 84.52 13.76 59.4 0 120.67-15.31 162.9-25.86 15.85-3.96 29.54-7.39 38.34-8.78 22.73-3.6 64.57-3.26 117.82.94 16.18 1.28 32.28-4.3 44.19-15.311 11.9-10.99 18.73-26.59 18.73-42.8V80.213c0-29.29-18.29-55.99-45.5-66.45zm3.5 335.28c0 4.6-1.86 8.84-5.23 11.96-3.39 3.12-7.78 4.65-12.38 4.28-39.94-3.15-94.51-5.81-127.7-.55-10.62 1.68-24.45 5.14-41.95 9.51-58.39 14.59-156.12 39.02-222.18 13.64-11.1-4.27-18.56-15.21-18.56-27.24V91.803c0-4.59 1.86-8.83 5.23-11.95 3.38-3.13 7.77-4.65 12.38-4.28 39.93 3.15 94.5 5.81 127.7.55 10.62-1.68 24.44-5.14 41.95-9.51 58.39-14.59 156.12-39.021 222.17-13.64 11.11 4.26 18.57 15.21 18.57 27.24v268.83z\"></path><path d=\"M259.091 297.086c-42.466 0-77.015-34.539-77.015-76.993s34.549-76.993 77.015-76.993 77.016 34.539 77.016 76.993-34.549 76.993-77.016 76.993zm0-111.986c-19.308 0-35.015 15.698-35.015 34.993s15.707 34.993 35.015 34.993 35.016-15.698 35.016-34.993-15.708-34.993-35.016-34.993z\"></path></g></svg> Payments</a> <a class=\"flex-none py-2 px-3 opacity-0 hover:opacity-100\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" class=\"fill-current w-4 h-4\"><path d=\"M255.8 40.8c-118.9 0-215.4 96.5-215.4 215.4s96.5 215.4 215.4 215.4 215.4-96.4 215.4-215.4S374.7 40.8 255.8 40.8zm94.5 237h-72.9v72.9c-.3 11.9-10.1 21.4-22.1 21.2-11.6-.2-20.9-9.6-21.2-21.2v-72.9h-72.9c-11.9 0-21.6-9.7-21.6-21.6s9.7-21.6 21.6-21.6h72.9v-72.9c.3-11.9 10.1-21.4 22.1-21.2 11.6.2 20.9 9.6 21.2 21.2v72.9h72.9c11.9 0 21.6 9.7 21.6 21.6 0 12-9.7 21.6-21.6 21.6z\"></path></svg></a></li><li class=\"mb-3 flex rounded hover:bg-primarydarker items-center\"><a href=\"#\" hx-get=\"/customer\" hx-target=\"#main_content\" hx-indicator=\"#spinner\" hx-swap=\"innerHTML\" hx-push-url=\"true\" class=\"grow flex items-center py-2 px-3 text-sm \"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 459.1 512\" class=\"fill-current w-6 h-4 mr-2\"><g id=\"Layer_1-2\"><path d=\"M382.4 412.6c11 0 20-9 20-20V20c0-11-9-20-20-20s-20 9-20 20v372.6c0 11 9 20 20 20zM77.7 412.6c11 0 20-9 20-20V256c0-11-9-20-20-20s-20 9-20 20v136.6c0 11 9 20 20 20zM229.1 412.6c11 0 20-9 20-20V133.9c0-11-9-20-20-20s-20 9-20 20v258.6c0 11 9 20 20 20zM439.1 472H20c-11 0-20 9-20 20s9 20 20 20h419.1c11 0 20-9 20-20s-9-20-20-20z\"></path></g></svg> Reports</a></li></ul></div></aside>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
