---
layout: 'okta'
page_title: 'Okta: okta_app_shared_credentials'
sidebar_current: 'docs-okta-resource-app-shared-credentials'
description: |-
    Creates a SWA shared credentials app.
---

# okta_app_shared_credentials

This resource allows you to create and configure SWA shared credentials app.

-> During an apply if there is change in `status` the app will first be
activated or deactivated in accordance with the `status` change. Then, all
other arguments that changed will be applied.

## Example Usage

```hcl
resource "okta_app_shared_credentials" "example" {
  label                            = "Example App"
  status                           = "ACTIVE"
  button_field                     = "btn-login"
  username_field                   = "txtbox-username"
  password_field                   = "txtbox-password"
  url                              = "https://example.com/login.html"
  redirect_url                     = "https://example.com/redirect_url"
  checkbox                         = "checkbox_red"
  user_name_template               = "user.firstName"
  user_name_template_type          = "CUSTOM"
  user_name_template_suffix        = "hello"
  shared_password                  = "sharedpass"
  shared_username                  = "sharedusername"
  accessibility_self_service       = true
  accessibility_error_redirect_url = "https://example.com/redirect_url_1"
  accessibility_login_redirect_url = "https://example.com/redirect_url_2"
  auto_submit_toolbar              = true
  hide_ios                         = true
}
```

## Argument Reference

The following arguments are supported:

- `accessibility_error_redirect_url` - (Optional) Custom error page URL.

- `accessibility_login_redirect_url` - (Optional) Custom login page for this application.

- `accessibility_self_service` - (Optional) Enable self-service. By default, it is `false`.

- `admin_note` - (Optional) Application notes for admins.

- `app_links_json` - (Optional) Displays specific appLinks for the app. The value for each application link should be boolean.

- `auto_submit_toolbar` - (Optional) Display auto submit toolbar.

- `button_field` - (Optional) CSS selector for the Sign-In button in the sign-in form.

- `checkbox` - (Optional) CSS selector for the checkbox.

- `enduser_note` - (Optional) Application notes for end users.

- `hide_ios` - (Optional) Do not display application icon on mobile app.

- `hide_web` - (Optional) Do not display application icon to users.

- `label` - (Required) The Application's display name.

- `logo` - (Optional) Local file path to the logo. The file must be in PNG, JPG, or GIF format, and less than 1 MB in size.

- `password_field` - (Optional) CSS selector for the Password field in the sign-in form.

- `preconfigured_app` - (Optional) name of application from the Okta Integration Network, if not included a custom app will be created.

- `redirect_url` - (Optional) Redirect URL. If going to the login page URL redirects to another page, then enter that URL here.

- `shared_password` - (Optional) Shared password, required for certain schemes.

- `shared_username` - (Optional) Shared username, required for certain schemes.

- `status` - (Optional) The status of the application, by default, it is `"ACTIVE"`.

- `url` - (Optional) The URL of the sign-in page for this app.

- `url_regex` - (Optional) A regular expression that further restricts url to the specified regular expression.

- `user_name_template` - (Optional) Username template. Default: `"${source.login}"`

- `user_name_template_push_status` - (Optional) Push username on update. Valid values: `"PUSH"` and `"DONT_PUSH"`.

- `user_name_template_suffix` - (Optional) Username template suffix.

- `user_name_template_type` - (Optional) Username template type. Default: `"BUILT_IN"`.

- `username_field` - (Optional) CSS selector for the username field.

## Attributes Reference

- `id` - ID of an app.

- `name` - Name assigned to the application by Okta.

- `sign_on_mode` - Sign-on mode of the application.

- `logo_url` - Direct link of application logo.

- `sign_on_mode` - Authentication mode of app.

## Timeouts

The `timeouts` block allows you to specify custom [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions: 

- `create` - Create timeout (default 1 hour).

- `update` - Update timeout (default 1 hour).

- `read` - Read timeout (default 1 hour).

## Import

Okta SWA Shared Credentials App can be imported via the Okta ID.

```
$ terraform import okta_app_shared_credentials.example &#60;app id&#62;
```
