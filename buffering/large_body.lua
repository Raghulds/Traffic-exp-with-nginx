wrk.method = "POST"
wrk.headers["Content-Type"] = "application/octet-stream"

local body = string.rep("A", 5 * 1024 * 1024) -- 5MB
wrk.body = body
