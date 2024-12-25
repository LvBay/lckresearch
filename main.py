import requests, openpyxl,time
from lxpy import lxheader
from multiprocessing.dummy import Pool

url_items = "https://lol.ps/api/match/{}/summary.json?summoner_id={}&region=kr"
url_timestamp = "https://lol.ps/api/summoner/{}/recent-matches.json?page={}&region=kr"








def get_response(url,sid,t):
    """
    发起请求
    :param url: 指定url
    :param sid: 指定用户sid
    :param t:指定页数/match_id
    :return:
    """
    User_agent = lxheader.get_ua()['user-agent']
    payload = {}
    headers = {
        'Cookie': '__cf_bm=3TImAIbtAvMh7_ITiYS10Rv0P1I9VPUfbuCwk4Ukh8I-1734680718-1.0.1.1-QFFSTp.rvAnzMRsZlpIrFZLspEEQ4vwqVDtEtUImVFdGiG5bPmoDc8Yyeq743Ht9R_uCJroZHyWgcRI7.4NmwQ',
        'User-Agent': User_agent
    }
    response = requests.request("GET", url.format(sid,t), headers=headers, data=payload)
    return  response.json()



def parse_match_id(sid,t):
    """
    解析获取match_id
    :param sid: 指定sid
    :param t:  指定页数（每页20条数据）
    :return:  match_id_list
    """
    json_data = get_response(url_timestamp,sid,t)
    match_id_list = []
    for i in json_data["data"]["recentMatches"]:
        match_id_list.append(i["match_id"])
    return match_id_list



def parse_items(match_id,sid):
    """
    获取items
    :param timestamp: timestamp
    :param match_id: match_id
    :return: itmes
    """
    json_data = get_response(url_items,match_id,sid)
    return json_data["data"]["participants"][f"{sid}"]["info"]["items"]

def parse_text(k):
    """
    请求获取中文名
    :param k: item
    :return: text
    """
    url = "https://lol.ps/api/info/item-info/{}"
    User_agent = lxheader.get_ua()['user-agent']
    payload = {}
    headers = {
        'Cookie': '__cf_bm=3TImAIbtAvMh7_ITiYS10Rv0P1I9VPUfbuCwk4Ukh8I-1734680718-1.0.1.1-QFFSTp.rvAnzMRsZlpIrFZLspEEQ4vwqVDtEtUImVFdGiG5bPmoDc8Yyeq743Ht9R_uCJroZHyWgcRI7.4NmwQ',
        'User-Agent': User_agent
    }
    response = requests.request("GET", url.format(k), headers=headers, data=payload)
    return response.json()["data"]["nameCn"]

def main(sid_name):
    # 初始化excelx表头
    sid,name = sid_name[0],sid_name[1]
    xlsx = openpyxl.Workbook()
    table = xlsx.active
    for a, b in zip([1, 2, 3, 4, 5, 6, 7], range(1, 8)):
        table.cell(1, b).value = a
    for t in range(1,4):
        time.sleep(2)
        match_id_list = parse_match_id(sid, t)
        for match_id,b in zip(match_id_list,range(1,21)):
            items = parse_items(match_id,sid)
            row = table.max_row
            for k,p in zip(items,range(1,8)):
                try:
                    table.cell(row + 1, p).value = parse_text(k)
                except:
                    pass
            print(f"写入{name}-{t}-{b}")
        xlsx.save(f'{name}.xlsx')


if __name__ == '__main__':
    sid_list = ["4txgTW5VnMA2NhHKwrQu1L-2exTQvUjmL_-ADqSBILH9po0","nenuR9JdGcGKe0mFNp7B6O3jOvSXCnnf91u69FpoWpcbReo",
                "DDkJH9AwqqBlNY3yNL7J5pbG68ulgeMmHsURk-ng4aah1Ks","tvy6jZOCuTXoLP8kRH3TPw-TadJ8zs0rQR7B545UkyD1-Ro",
                "OsxLzBfUmY03Tdp-4HWnv4Fw0HosMq2v2eB6STFX4BywnKnU","c4tXq5cTaD7Vo4exEO0Fpwqel3MnoP21Q77bE7ONFmqMtTu8EVuwq97Jow",
                "m4-vweCLnO2BetTyKXpp96rU6SlxJc1Y9LHMCy7obn-byWw","mcTBMe2P-t66RwHx514490RpYyezaPNXzlXVHE-Dm2Dzu0v-",
                "t2juS3_APvngM9uWTOepXvP1_5Tg0dLaP1gxDQ38Ssc3pABPzg-FzoMm8A","VB7kbMWzlR9_aUEANWLWkkeP0Hz_RkgDDJz-uCgWpniUtp-K"]
    name_list = ["민철이여친구함#0415","칼과 창 방패#KR1","Yondaime#Luo","Morgan#5358","Happy#day12","Ivory#1012(1)","Ivory#1012(2)","Raptor#KR123",
                 "DANDAN心魅かれてく#solo","클래식좀들어라#KR2"]
    list = []
    for sid,name in zip(sid_list,name_list):
        list.append([sid,name])
    pool = Pool(1)  # 指定线程数
    pool.map(main,list)