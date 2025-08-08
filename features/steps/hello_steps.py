from behave import given, when, then
import requests, json

@given('the API base url is "{base}"')
def _(ctx, base):
    ctx.url = base

@when('I GET "{path}"')
def _(ctx, path):
    ctx.resp = requests.get(ctx.url + path, timeout=3)

@then('the response code is {code:d}')
def _(ctx, code):
    assert ctx.resp.status_code == code, ctx.resp.text

@then('the body contains "{text}"')
def _(ctx, text):
    assert text in ctx.resp.text
