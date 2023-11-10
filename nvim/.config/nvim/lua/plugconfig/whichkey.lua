local wk = require("which-key")
wk.setup {
  icons = {
    group = ''
  }
}

wk.register({
  ["w"] = {
    name = "Windows", -- optional group name
  },
  ["<leader>t"] = {
    name = "Tabs",
  },
  ["<leaver>c"] = {
    name = "CMake"
  }
})
