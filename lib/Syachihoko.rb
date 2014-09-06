# -*- coding: utf-8 -*-
class Syachihoko
  def syachihoko()
  end

  # アクセスログのファイルを引数にアクセス数を取得する
  def collect(filename)
    access_count = {};
    # ファイルを開いてアクセスパスごとに数を合計する
    open(filename) do |f|
      while line = f.gets
        logs = line.split("\s");
        path = logs[6]
        if access_count.has_key?(path) then
          access_count[path] += 1
        else
          access_count[path] = 1
        end
      end
    end
    # アクセス数で並び替えて戻す
    return access_count.sort{|(k1, v1), (k2, v2)|
      v2 <=> v1
    }
  end
end
