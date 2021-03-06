
context = ChefDK::Generator.context
cookbook_dir = File.join(context.cookbook_root, context.cookbook_name)

silence_chef_formatter unless context.verbose

generator_desc("Ensuring correct cookbook file content")

# cookbook root dir
directory cookbook_dir

# metadata.rb
template "#{cookbook_dir}/metadata.rb" do
  helpers(ChefDK::Generator::TemplateHelper)
  action :create_if_missing
end

# README
template "#{cookbook_dir}/README.md" do
  helpers(ChefDK::Generator::TemplateHelper)
  action :create_if_missing
end

# chefignore
cookbook_file "#{cookbook_dir}/chefignore"

if context.use_berkshelf

  # Berks
  cookbook_file "#{cookbook_dir}/Berksfile" do
    action :create_if_missing
  end
else

  # Policyfile
  template "#{cookbook_dir}/Policyfile.rb" do
    source "Policyfile.rb.erb"
    helpers(ChefDK::Generator::TemplateHelper)
  end

end

# Chefspec
directory "#{cookbook_dir}/spec/unit/recipes" do
  recursive true
end

cookbook_file "#{cookbook_dir}/spec/spec_helper.rb" do

  if context.use_berkshelf
    source "spec_helper.rb"
  else
    source "spec_helper_policyfile.rb"
  end

  action :create_if_missing
end

template "#{cookbook_dir}/spec/unit/recipes/default_spec.rb" do
  source "recipe_spec.rb.erb"
  helpers(ChefDK::Generator::TemplateHelper)
  action :create_if_missing
end

# Recipes

directory "#{cookbook_dir}/recipes"

template "#{cookbook_dir}/recipes/default.rb" do
  source "recipe.rb.erb"
  helpers(ChefDK::Generator::TemplateHelper)
  action :create_if_missing
end

# git
if context.have_git
  unless context.skip_git_init

    generator_desc("Committing cookbook files to git")

    execute("initialize-git") do
      command("git init .")
      cwd cookbook_dir
    end

  end

  cookbook_file "#{cookbook_dir}/.gitignore" do
    source "gitignore"
  end

  unless context.skip_git_init

    execute("git-add-new-files") do
      command("git add .")
      cwd cookbook_dir
    end

    execute("git-commit-new-files") do
      command("git commit -m \"Add generated cookbook content\"")
      cwd cookbook_dir
    end
  end
end

if context.enable_delivery

  include_recipe "::build_cookbook"

end
