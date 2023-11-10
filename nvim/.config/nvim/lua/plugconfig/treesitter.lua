local configs = require("nvim-treesitter.configs")

configs.setup({
  ensure_installed = { "c", "lua", "vim", "vimdoc", "query", "cpp", "go", "javascript", "html", "d", "rust" },
  sync_install = false,
  auto_install = false,
  highlight = { enable = true },
  indent = { enable = true }, 
})
