What Casbin does:

* enforce the policy in the classic {subject, object, action} form or a customized form as you defined, both allow and deny authorizations are supported.
* handle the storage of the access control model and its policy.
* manage the role-user mappings and role-role mappings (aka role hierarchy in RBAC).
* support built-in superuser like root or administrator. A superuser can do anything without explict permissions.
* multiple built-in operators to support the rule matching. For example, keyMatch can map a resource key /foo/bar to the pattern /foo*.
