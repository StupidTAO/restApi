package model

import "errors"

//TODO: token model operation

func (tokenId TokenId) String() (string, error) {
	switch tokenId {
	case AOA:
		//todo
		return "AOA_ar_3", nil
	case BAT:
		return "BAT_ar_3", nil

	case BNB:
		return "BNB_ar_3", nil
	case CRO:
		return "CRO_ar_3", nil
	case DAI:
		return "DAI_ar_3", nil
	case ENJ:
		return "ENJ_ar_3", nil
	case GNT:
		return "GNT_ar_3", nil
	case HOT:
		return "HOT_ar_3", nil
	case HT:
		return "HT", nil
	case INB:
		return "INB", nil
	case KCS:
		return "KCS", nil
	case LAMB:
		return "LAMB", nil
	case LEO:
		return "LEO", nil
	case LINK:
		return "LINK", nil
	case MCO:
		return "MCO", nil
	case MKR:
		return "MKR", nil
	case NEXO:
		return "NEXO_ar_3", nil
	case NPXS:
		return "NPXS_ar_3", nil
	case PAX:
		return "PAX_ar_3", nil
	case QNT:
		return "QNT_ar_3", nil
	case REP:
		return "REP_ar_3", nil
	case SNT:
		return "SNT_ar_3", nil
	case SNX:
		return "SNX_ar_3", nil
	case TUSD:
		return "TUSD_ar_3", nil
	case USDC:
		return "USDC_ar_3", nil
	case USDT:
		return "USDT_ar_3", nil
	case ZB:
		return "ZB_ar_3", nil
	case ZRX:
		return "ZRX_ar_3", nil
	default:
		return "", errors.New("invalid tokenId")
		
	}
}

func StrFormatTokenId(tokenStr string) (TokenId, error){
	switch tokenStr {
	case "AOA_ar_3":
		return AOA, nil
	case "BAT_ar_3":
		return BAT, nil
	case "BNB_ar_3":
		return BNB, nil
	case "CRO_ar_3":
		return CRO, nil
	case "ENJ_ar_3":
		return ENJ, nil
	case "GNT_ar_3":
		return GNT, nil
	case "HOT_ar_3":
		return HOT, nil
	case "HT_ar_3":
		return HT, nil
	case "INB_ar_3":
		return INB, nil
	case "KCS_ar_3":
		return KCS, nil
	case "LAMB_ar_3":
		return LAMB, nil
	case "LEO_ar_3":
		return LEO, nil
	case "LINK_ar_3":
		return LINK, nil
	case "MCO_ar_3":
		return MCO, nil
	case "MKR_ar_3":
		return MKR, nil
	case "NEXO_ar_3":
		return NEXO, nil
	case "NPXS_ar_3":
		return NPXS, nil
	case "PAX_ar_3":
		return PAX, nil
	case "QNT_ar_3":
		return QNT, nil
	case "REP_ar_3":
		return REP, nil
	case "SNT_ar_3":
		return SNT, nil
	case "SNX_ar_3":
		return SNX, nil
	case "TUSD_ar_3":
		return TUSD, nil
	case "USDT_ar_3":
		return USDT, nil
	case "ZB_ar_3":
		return ZB, nil
	case "ZRX_ar_3":
		return ZRX, nil
	}

	return INVALID, errors.New("invalid string token name")
}